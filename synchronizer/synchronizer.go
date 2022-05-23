package synchronizer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"

	"github.com/niftyConnect/order-synchronizer/adaptor"
	localcmm "github.com/niftyConnect/order-synchronizer/common"
	"github.com/niftyConnect/order-synchronizer/config"
	"github.com/niftyConnect/order-synchronizer/database"
)

type Synchronizer struct {
	blockchain  string
	cfg         *config.ChainDetail
	db          *gorm.DB
	adaptorInst *adaptor.Adaptor
}

func NewSynchronizer(db *gorm.DB, cfg *config.ChainDetail, blockchain string) (*Synchronizer, error) {
	adaptorInst, err := adaptor.NewAdaptor(cfg, blockchain)
	if err != nil {
		return nil, err
	}
	return &Synchronizer{
		blockchain:  blockchain,
		cfg:         cfg,
		db:          db,
		adaptorInst: adaptorInst,
	}, nil
}

func (syncInst *Synchronizer) Start() {
	go syncInst.fetchDaemon()
	go syncInst.expireOrdersDeamon()
	go syncInst.pruneDaemon()
}

func (syncInst *Synchronizer) fetchDaemon() {
	defer func() {
		if r := recover(); r != nil {
			localcmm.Logger.Errorf("------------------------------------------")
			localcmm.Logger.Errorf("Recovered from panic %v", r)
			localcmm.Logger.Errorf("------------------------------------------")
		}
	}()

	curBlockLog, err := syncInst.getCurrentBlockLog()
	if err != nil {
		panic(fmt.Sprintf("get current block log from db error: %s", err.Error()))
	}

	latestHeader, err := syncInst.adaptorInst.GetLatestHeader(context.Background())
	if err != nil {
		panic(fmt.Sprintf("get latest block header error: %s", err.Error()))
	}

	latestConfirmedHeader, err := syncInst.adaptorInst.GetHeaderByNumber(context.Background(), big.NewInt(latestHeader.Number.Int64()-syncInst.cfg.ConfirmBlocks))
	if err != nil {
		panic(fmt.Sprintf("get latest block header error: %s", err.Error()))
	}

	localcmm.Logger.Infof("latest height in DB %d, latest %s height %d", curBlockLog.Height, syncInst.blockchain, latestConfirmedHeader.Number.Int64())

	if curBlockLog == nil || latestConfirmedHeader.Number.Int64() > curBlockLog.Height+syncInst.cfg.HeightStep {
		localcmm.Logger.Infof("Extract history events")
		events, err := syncInst.adaptorInst.ExtractEvents(syncInst.cfg.HeightStep, syncInst.cfg.StartHeight, latestConfirmedHeader.Number.Int64())
		if err != nil {
			panic(fmt.Sprintf("extract events from %d to %d, error: %s", syncInst.cfg.StartHeight, uint64(latestConfirmedHeader.Number.Int64()), err.Error()))
		}
		nextBlockLog := database.BlockLog{
			Blockchain: syncInst.blockchain,
			BlockHash:  latestConfirmedHeader.Hash().String(),
			ParentHash: "",
			Height:     latestConfirmedHeader.Number.Int64(),
			BlockTime:  int64(latestConfirmedHeader.Time),
		}
		curBlockLog = &nextBlockLog
		err = syncInst.saveBlockAndTxEvents(nextBlockLog, events)
		if err != nil {
			panic(err)
		}

		err = syncInst.analysisOrder(latestConfirmedHeader.Number.Int64())
		if err != nil {
			panic(err)
		}
	}

	localcmm.Logger.Infof("Start fetchDaemon")
	for {
		func() {
			localcmm.Logger.Infof("try to fetch %s block, height=%d", syncInst.blockchain, curBlockLog.Height+1)
			newHeight, newBlockHash, err := syncInst.fetchBlock(curBlockLog.Height, curBlockLog.Height+1, curBlockLog.BlockHash)
			if err != nil {
				localcmm.Logger.Debugf("fetch %s block error, err=%s", syncInst.blockchain, err.Error())
				time.Sleep(time.Duration(syncInst.cfg.SleepWaitSecond) * time.Second)
			} else {
				curBlockLog.Height = newHeight
				curBlockLog.BlockHash = newBlockHash.String()

				err = syncInst.analysisOrder(newHeight - syncInst.cfg.ConfirmBlocks)
				if err != nil {
					panic(err)
				}
			}
		}()
	}
}

func (syncInst *Synchronizer) expireOrdersDeamon() {
	defer func() {
		if r := recover(); r != nil {
			localcmm.Logger.Errorf("------------------------------------------")
			localcmm.Logger.Errorf("Recovered from panic %v", r)
			localcmm.Logger.Errorf("------------------------------------------")
		}
	}()
	for {
		time.Sleep(expireInterval)

		tx := syncInst.db.Begin()
		if err := tx.Error; err != nil {
			localcmm.Logger.Errorf("%s", err.Error())
			continue
		}

		currentTime := time.Now().Unix()
		if err := tx.Model(database.NiftyConnectOrder{}).Where("blockchain = ? and expiration_time < ? and is_expired = ?", syncInst.blockchain, currentTime, false).
			Updates(map[string]interface{}{
				"is_expired":  true,
				"update_time": time.Now().Unix(),
			}).Error; err != nil {
			tx.Rollback()
			localcmm.Logger.Errorf("%s", err.Error())
			continue
		}

		if err := tx.Commit().Error; err != nil {
			localcmm.Logger.Errorf("%s", err.Error())
			continue
		}
	}
}

func (syncInst *Synchronizer) pruneDaemon() {
	defer func() {
		if r := recover(); r != nil {
			localcmm.Logger.Errorf("------------------------------------------")
			localcmm.Logger.Errorf("Recovered from panic %v", r)
			localcmm.Logger.Errorf("------------------------------------------")
		}
	}()
	for {
		curBlockLog, err := syncInst.getCurrentBlockLog()
		if err != nil {
			localcmm.Logger.Errorf("get current block log error, err=%s", err.Error())
			time.Sleep(pruneInterval)
			continue
		}
		err = syncInst.db.Where("blockchain = ? and height < ?",
			syncInst.blockchain,
			curBlockLog.Height-pruneHeightGap).
			Delete(database.BlockLog{}).Error
		if err != nil {
			localcmm.Logger.Infof("prune block logs error, err=%s", err.Error())
		}
		time.Sleep(pruneInterval)
	}
}

func (syncInst *Synchronizer) getCurrentBlockLog() (*database.BlockLog, error) {
	blockLog := database.BlockLog{}
	err := syncInst.db.Model(database.BlockLog{}).Where("blockchain = ?", syncInst.blockchain).Order("height desc").First(&blockLog).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &blockLog, nil
}

func (syncInst *Synchronizer) analysisOrder(latestEventHeight int64) error {
	isOrderAnalysisStatusExist := true
	syncStatus := database.OrderAnalysisStatus{}
	err := syncInst.db.Model(database.OrderAnalysisStatus{}).Where("blockchain = ?", syncInst.blockchain).First(&syncStatus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	} else if err == gorm.ErrRecordNotFound {
		isOrderAnalysisStatusExist = false
	}
	latestAnalysisHeight := syncStatus.LatestAnalysisHeight

	var orderApprovedPartOneList []database.OrderApprovedPartOne
	if err := syncInst.db.Model(database.OrderApprovedPartOne{}).
		Where("blockchain = ? and height > ? and height <= ?", syncInst.blockchain, latestAnalysisHeight, latestEventHeight).
		Find(&orderApprovedPartOneList).Error; err != nil {
		return err
	}
	var orderApprovedPartTwoList []database.OrderApprovedPartTwo
	if err := syncInst.db.Model(database.OrderApprovedPartTwo{}).
		Where("blockchain = ? and height > ? and height <= ?", syncInst.blockchain, latestAnalysisHeight, latestEventHeight).
		Find(&orderApprovedPartTwoList).Error; err != nil {
		return err
	}
	var orderCancelledList []database.OrderCancelled
	if err := syncInst.db.Model(database.OrderCancelled{}).
		Where("blockchain = ? and height > ? and height <= ?", syncInst.blockchain, latestAnalysisHeight, latestEventHeight).
		Find(&orderCancelledList).Error; err != nil {
		return err
	}
	var ordersMatchedList []database.OrdersMatched
	if err := syncInst.db.Model(database.OrdersMatched{}).
		Where("blockchain = ? and height > ? and height <= ?", syncInst.blockchain, latestAnalysisHeight, latestEventHeight).
		Find(&ordersMatchedList).Error; err != nil {
		return err
	}
	var nonceIncrementedList []database.NonceIncremented
	if err := syncInst.db.Model(database.NonceIncremented{}).
		Where("blockchain = ? and height > ? and height <= ?", syncInst.blockchain, latestAnalysisHeight, latestEventHeight).
		Find(&nonceIncrementedList).Error; err != nil {
		return err
	}

	newCreateOrders := make([]database.NiftyConnectOrder, 0, len(orderApprovedPartOneList))
	for _, eventPartOne := range orderApprovedPartOneList {
		var eventPartTwo database.OrderApprovedPartTwo
		foundPartTwo := false
		for _, eventPartTwoItem := range orderApprovedPartTwoList {
			if eventPartOne.Hash == eventPartTwoItem.Hash {
				eventPartTwo = eventPartTwoItem
				foundPartTwo = true
				break
			}
		}
		if !foundPartTwo {
			return fmt.Errorf("failed to find order approved part two event, order hash %s", eventPartOne.Hash)
		}

		newCreateOrders = append(newCreateOrders, database.NiftyConnectOrder{
			Blockchain:               syncInst.blockchain,
			Height:                   eventPartOne.Height,
			OrderHash:                eventPartOne.Hash,
			TxHash:                   eventPartOne.TxHash,
			Exchange:                 eventPartOne.Exchange,
			Maker:                    eventPartOne.Maker,
			Taker:                    eventPartOne.Taker,
			MakerRelayerFeeRecipient: eventPartOne.MakerRelayerFeeRecipient,
			TakerRelayerFeeRecipient: zeroAddress.String(),
			Side:                     eventPartOne.Side,
			SaleKind:                 eventPartOne.SaleKind,
			NftAddress:               eventPartOne.NftAddress,
			TokenId:                  eventPartOne.TokenId,
			IpfsHash:                 eventPartOne.IpfsHash,
			Calldata:                 eventPartTwo.Calldata,
			ReplacementPattern:       eventPartTwo.ReplacementPattern,
			StaticTarget:             eventPartTwo.StaticTarget,
			StaticExtradata:          eventPartTwo.StaticExtradata,
			PaymentToken:             eventPartTwo.PaymentToken,
			OrderPrice:               eventPartTwo.BasePrice,
			Extra:                    eventPartTwo.Extra,
			ListingTime:              eventPartTwo.ListingTime,
			ExpirationTime:           eventPartTwo.ExpirationTime,
			Salt:                     eventPartTwo.Salt,
			IsCancelled:              false,
			IsExpired:                false,
			IsFinalized:              false,
			CreateTime:               time.Now().Unix(),
			UpdateTime:               time.Now().Unix(),
		})
	}

	for idx, newCreateOrder := range newCreateOrders {
		for _, orderCancelled := range orderCancelledList {
			if newCreateOrder.OrderHash == orderCancelled.Hash {
				newCreateOrders[idx].IsCancelled = true
			}
		}
		for _, ordersMatched := range ordersMatchedList {
			if newCreateOrder.OrderHash == ordersMatched.BuyHash {
				newCreateOrders[idx].IsFinalized = true
			}
			if newCreateOrder.OrderHash == ordersMatched.SellHash {
				newCreateOrders[idx].IsFinalized = true
			}
		}
		for _, nonceIncremented := range nonceIncrementedList {
			if newCreateOrder.Maker == nonceIncremented.Maker &&
				newCreateOrder.Height < nonceIncremented.Height || newCreateOrder.Height == nonceIncremented.Height && newCreateOrder.Position < nonceIncremented.Position {
				newCreateOrders[idx].IsCancelled = true
			}
		}
	}

	cancelOrderHashList := make([]string, 0, len(orderCancelledList))
	for _, orderCancelled := range orderCancelledList {
		cancelOrderHashList = append(cancelOrderHashList, orderCancelled.Hash)
	}

	matchedOrderHashList := make([]string, 0, 2*len(ordersMatchedList))
	for _, ordersMatched := range ordersMatchedList {
		matchedOrderHashList = append(matchedOrderHashList, ordersMatched.BuyHash)
		matchedOrderHashList = append(matchedOrderHashList, ordersMatched.SellHash)
	}

	nonceIncrementedMakerList := make([]string, 0, len(nonceIncrementedList))
	nonceIncrementedMakerMap := make(map[string]bool)
	for _, nonceIncremented := range nonceIncrementedList {
		if nonceIncrementedMakerMap[nonceIncremented.Maker] {
			continue
		}
		nonceIncrementedMakerList = append(nonceIncrementedMakerList, nonceIncremented.Maker)
	}

	tx := syncInst.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	for _, newCreateOrder := range newCreateOrders {
		if err := tx.Create(&newCreateOrder).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Model(database.NiftyConnectOrder{}).Where("blockchain = ? and order_hash in (?)", syncInst.blockchain, cancelOrderHashList).
		Updates(map[string]interface{}{
			"is_cancelled": true,
			"update_time":  time.Now().Unix(),
		}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(database.NiftyConnectOrder{}).Where("blockchain = ? and order_hash in (?)", syncInst.blockchain, matchedOrderHashList).
		Updates(map[string]interface{}{
			"is_finalized": true,
			"update_time":  time.Now().Unix(),
		}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(database.NiftyConnectOrder{}).Where("blockchain = ? and maker in (?)", syncInst.blockchain, nonceIncrementedMakerList).
		Updates(map[string]interface{}{
			"is_cancelled": true,
			"update_time":  time.Now().Unix(),
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if !isOrderAnalysisStatusExist {
		orderAnalysisStatus := database.OrderAnalysisStatus{
			Blockchain:           syncInst.blockchain,
			LatestAnalysisHeight: latestEventHeight,
			CreateTime:           time.Now().Unix(),
			UpdateTime:           time.Now().Unix(),
		}
		if err := tx.Create(&orderAnalysisStatus).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := tx.Model(database.OrderAnalysisStatus{}).Where("blockchain = ?", syncInst.blockchain).
			Updates(map[string]interface{}{
				"latest_analysis_height": latestEventHeight,
				"update_time":            time.Now().Unix(),
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	localcmm.Logger.Infof("Analysis order to height %d", latestEventHeight)
	return tx.Commit().Error
}

func (syncInst *Synchronizer) fetchBlock(curHeight, nextHeight int64, curBlockHash string) (int64, common.Hash, error) {
	header, err := syncInst.adaptorInst.GetHeaderByNumber(context.Background(), big.NewInt(nextHeight))
	if err != nil {
		return 0, common.Hash{}, err
	}

	parentHash := header.ParentHash.String()
	if curHeight != 0 && parentHash != curBlockHash {
		localcmm.Logger.Infof("blockchain %s fork detected, delete previous block and event, current height %d, curBlockHash %s, new block height %d, new block hash %s, new block parent hash %s",
			syncInst.blockchain, curHeight, curBlockHash, nextHeight, header.Hash().String(), parentHash)
		previousBlock, err := syncInst.getCurrentBlockLog()
		if err != nil {
			return 0, common.Hash{}, err
		}
		return previousBlock.Height, common.HexToHash(previousBlock.BlockHash), syncInst.deleteBlockAndTxEvents(curHeight)
	} else {
		nextBlockLog := database.BlockLog{
			Blockchain: syncInst.blockchain,
			BlockHash:  header.Hash().String(),
			ParentHash: parentHash,
			Height:     header.Number.Int64(),
			BlockTime:  int64(header.Time),
		}

		events, err := syncInst.adaptorInst.GetNiftyConnectEvents(header)
		if err != nil {
			return 0, common.Hash{}, err
		}

		err = syncInst.saveBlockAndTxEvents(nextBlockLog, events)
		if err != nil {
			return 0, common.Hash{}, err
		}
	}
	return header.Number.Int64(), header.Hash(), nil
}

func (syncInst *Synchronizer) deleteBlockAndTxEvents(height int64) error {
	tx := syncInst.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("height = ?", height).Delete(database.BlockLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("blockchain = ? and height = ?", syncInst.blockchain, height).Delete(database.BlockLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("blockchain = ? and height = ?", syncInst.blockchain, height).Delete(database.OrderApprovedPartOne{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("blockchain = ? and height = ?", syncInst.blockchain, height).Delete(database.OrderApprovedPartTwo{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("blockchain = ? and height = ?", syncInst.blockchain, height).Delete(database.OrderCancelled{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("blockchain = ? and height = ?", syncInst.blockchain, height).Delete(database.OrdersMatched{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("blockchain = ? and height = ?", syncInst.blockchain, height).Delete(database.NonceIncremented{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (syncInst *Synchronizer) saveBlockAndTxEvents(blockLog database.BlockLog, packages []interface{}) error {
	tx := syncInst.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&blockLog).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, ev := range packages {
		switch ev.(type) {
		case database.OrderApprovedPartOne:
			event := ev.(database.OrderApprovedPartOne)
			if err := tx.Create(&event).Error; err != nil {
				tx.Rollback()
				return err
			}
		case database.OrderApprovedPartTwo:
			event := ev.(database.OrderApprovedPartTwo)
			if err := tx.Create(&event).Error; err != nil {
				tx.Rollback()
				return err
			}
		case database.OrderCancelled:
			event := ev.(database.OrderCancelled)
			if err := tx.Create(&event).Error; err != nil {
				tx.Rollback()
				return err
			}
		case database.OrdersMatched:
			event := ev.(database.OrdersMatched)
			if err := tx.Create(&event).Error; err != nil {
				tx.Rollback()
				return err
			}
		case database.NonceIncremented:
			event := ev.(database.NonceIncremented)
			if err := tx.Create(&event).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
