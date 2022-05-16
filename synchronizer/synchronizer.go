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
	blockchain    string
	cfg           *config.ChainDetail
	db            *gorm.DB
	adaptorInst   *adaptor.Adaptor

	currentHeight int64
}

func NewSynchronizer(db *gorm.DB, cfg *config.Config, blockchain string) (*Synchronizer, error) {
	adaptorInst, err := adaptor.NewAdaptor(cfg, blockchain)
	if err != nil {
		return nil, err
	}
	return &Synchronizer{
		blockchain:  blockchain,
		cfg:         cfg.ChainConfig.ChainDetails[blockchain],
		db:          db,
		adaptorInst: adaptorInst,
	}, nil
}

func (syncInst *Synchronizer) Start() {
	go syncInst.fetchDaemon(syncInst.cfg.StartHeight)
	go syncInst.pruneDaemon()
}

func (syncInst *Synchronizer) fetchDaemon(startHeight int64) {
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

	if curBlockLog == nil || latestHeader.Number.Int64() > curBlockLog.Height+syncInst.cfg.HeightStep {
		events, err := syncInst.adaptorInst.ExtractEvents(syncInst.cfg.HeightStep, syncInst.cfg.StartHeight, latestHeader.Number.Int64())
		if err != nil {
			panic(fmt.Sprintf("extract events from %d to %d, error: %s", syncInst.cfg.StartHeight, uint64(latestHeader.Number.Int64()), err.Error()))
		}
		nextBlockLog := database.BlockLog{
			Chain:      syncInst.blockchain,
			BlockHash:  latestHeader.Hash().String(),
			ParentHash: "",
			Height:     latestHeader.Number.Int64(),
			BlockTime:  int64(latestHeader.Time),
		}
		err = syncInst.saveBlockAndTxEvents(nextBlockLog, events)
		if err != nil {
			panic(err)
		}
	}

	for {
		func() {

			syncInst.currentHeight = curBlockLog.Height

			nextHeight := curBlockLog.Height + 1
			if curBlockLog.Height == 0 && startHeight != 0 {
				nextHeight = startHeight
			}

			localcmm.Logger.Debugf("fetch %s block, height=%d", syncInst.blockchain, nextHeight)
			newHeight, newBlockHash, err := syncInst.fetchBlock(curBlockLog.Height, nextHeight, curBlockLog.BlockHash)
			if err != nil {
				localcmm.Logger.Debugf("fetch %s block error, err=%s", syncInst.blockchain, err.Error())
				time.Sleep(time.Duration(syncInst.cfg.SleepWaitSecond) * time.Second)
			} else {
				curBlockLog.Height = newHeight
				curBlockLog.BlockHash = newBlockHash.String()
			}
		}()
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
			time.Sleep(observerPruneInterval)
			continue
		}
		err = syncInst.db.Where("chain = ? and height < ?",
			syncInst.blockchain,
			curBlockLog.Height-pruneHeightGap).
			Delete(database.BlockLog{}).Error
		if err != nil {
			localcmm.Logger.Infof("prune block logs error, err=%s", err.Error())
		}
		time.Sleep(observerPruneInterval)
	}
}

func (syncInst *Synchronizer) getCurrentBlockLog() (*database.BlockLog, error) {
	blockLog := database.BlockLog{}
	err := syncInst.db.Model(database.BlockLog{}).Where("chain = ?", syncInst.blockchain).Order("height desc").First(&blockLog).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &blockLog, nil
}

func (syncInst *Synchronizer) fetchBlock(curHeight, nextHeight int64, curBlockHash string) (int64, common.Hash, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	header, err := syncInst.adaptorInst.GetHeaderByNumber(ctxWithTimeout, big.NewInt(int64(nextHeight)))
	if err != nil {
		return 0, common.Hash{}, err
	}

	parentHash := header.ParentHash.String()
	if curHeight != 0 && parentHash != curBlockHash {
		localcmm.Logger.Infof("blockchain %s fork detected, delete previous block and event, height %d, parentHash %s, curBlockHash %s",
			syncInst.blockchain, curHeight, parentHash, curBlockHash)
		previousBlock, err := syncInst.getCurrentBlockLog()
		if err != nil {
			return 0, common.Hash{}, err
		}
		return previousBlock.Height, common.HexToHash(previousBlock.BlockHash), syncInst.deleteBlockAndTxEvents(curHeight)
	} else {
		nextBlockLog := database.BlockLog{
			Chain:      syncInst.blockchain,
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

	for _, pack := range packages {
		if err := tx.Create(&pack).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
