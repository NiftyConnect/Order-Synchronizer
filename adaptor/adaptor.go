package adaptor

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/niftyConnect/order-synchronizer/adaptor/contracts"
	localcmm "github.com/niftyConnect/order-synchronizer/common"
	"github.com/niftyConnect/order-synchronizer/config"
)

var (
	zeroAddress = common.Address{}
)

type Adaptor struct {
	blockchain              string
	cfg                     *config.ChainDetail
	niftyConnectExchangeABI abi.ABI
	niftyConnectAddress     common.Address
	clientMgr               *ClientMgr
}

func NewAdaptor(cfg *config.ChainDetail, blockchain string) (*Adaptor, error) {
	clientMgr, err := NewClientMgr(cfg, blockchain)
	if err != nil {
		return nil, err
	}

	niftyConnectExchangeABIV, err := abi.JSON(strings.NewReader(contracts.NiftyconnectexchangeABI))
	if err != nil {
		return nil, fmt.Errorf("marshal abi error")
	}
	return &Adaptor{
		blockchain:              blockchain,
		cfg:                     cfg,
		clientMgr:               clientMgr,
		niftyConnectExchangeABI: niftyConnectExchangeABIV,
		niftyConnectAddress:     cfg.ExchangeContract,
	}, nil
}

func (adaptorInst *Adaptor) GetLatestHeader(ctx context.Context) (*types.Header, error) {
	return adaptorInst.clientMgr.GetClient().HeaderByNumber(ctx, nil)
}

func (adaptorInst *Adaptor) GetHeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return adaptorInst.clientMgr.GetClient().HeaderByNumber(ctx, number)
}

func (adaptorInst *Adaptor) GetNiftyConnectEvents(header *types.Header) ([]interface{}, error) {
	return adaptorInst.getNiftyConnectEventAtBlock(header)
}

func (adaptorInst *Adaptor) getNiftyConnectEventAtBlock(header *types.Header) ([]interface{}, error) {
	topics := [][]common.Hash{{
		adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartOneEvent].ID,
		adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartTwoEvent].ID,
		adaptorInst.niftyConnectExchangeABI.Events[orderCancelledEvent].ID,
		adaptorInst.niftyConnectExchangeABI.Events[ordersMatchedEvent].ID,
		adaptorInst.niftyConnectExchangeABI.Events[nonceIncrementedEvent].ID,
	}}

	blockHash := header.Hash()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Second)
	defer cancel()

	logs, err := adaptorInst.clientMgr.GetClient().FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []common.Address{adaptorInst.niftyConnectAddress},
	})
	if err != nil {
		return nil, err
	}

	return adaptorInst.parseEventsFromLogs(logs)
}

func (adaptorInst *Adaptor) getNiftyConnectEventFromBlockRange(fromHeight, toHeight int64) ([]interface{}, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Second)
	defer cancel()

	topics := [][]common.Hash{{
		adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartOneEvent].ID,
		adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartTwoEvent].ID,
		adaptorInst.niftyConnectExchangeABI.Events[orderCancelledEvent].ID,
		adaptorInst.niftyConnectExchangeABI.Events[ordersMatchedEvent].ID,
		adaptorInst.niftyConnectExchangeABI.Events[nonceIncrementedEvent].ID,
	}}

	logs, err := adaptorInst.clientMgr.GetClient().FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		FromBlock: big.NewInt(fromHeight),
		ToBlock:   big.NewInt(toHeight),
		Topics:    topics,
		Addresses: []common.Address{adaptorInst.niftyConnectAddress},
	})
	if err != nil {
		return nil, err
	}

	return adaptorInst.parseEventsFromLogs(logs)
}

func (adaptorInst *Adaptor) ExtractEvents(initHeightStep int64, latestHeight, toHeight int64) ([]interface{}, error) {
	heightStep := initHeightStep
	startHeight := latestHeight + 1
	endHeight := startHeight + heightStep - 1
	result := make([]interface{}, 0)

	localcmm.Logger.Infof("Blockchain %s, delta filter event from %d to %d, initHeightStep %d",
		adaptorInst.blockchain, latestHeight, toHeight, initHeightStep)
	for startHeight <= toHeight {
		if endHeight > toHeight {
			endHeight = toHeight
		}
		localcmm.Logger.Infof("Blockchain %s, filter event between %d, %d, height step %d",
			adaptorInst.blockchain, startHeight, endHeight, heightStep)
	retry:
		events, err := func() (niftyConnectEvents []interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf(fmt.Sprintf("%v", r))
				}
			}()
			niftyConnectEvents, err = adaptorInst.getNiftyConnectEventFromBlockRange(startHeight, endHeight)
			if err != nil {
				return nil, fmt.Errorf("blockchain %s, extract NiftyConnect event err: %s", adaptorInst.blockchain, err)
			}
			return niftyConnectEvents, nil
		}()
		if err != nil && strings.Contains(err.Error(), "query returned more than 10000 results") {
			heightStep = heightStep * 4 / 5
			endHeight = startHeight + heightStep - 1
			localcmm.Logger.Errorf("reduce height step to %d, and sleep for a while and retry", heightStep)
			time.Sleep(5 * time.Second)
			goto retry
		} else if err != nil {
			return nil, err
		}
		result = append(result, events...)
		startHeight = endHeight + 1
		endHeight = startHeight + heightStep - 1
	}
	localcmm.Logger.Infof("Complete syncing recent events")
	return result, nil
}

func (adaptorInst *Adaptor) parseEventsFromLogs(logs []types.Log) ([]interface{}, error) {
	niftyConnectContract := bind.NewBoundContract(zeroAddress, adaptorInst.niftyConnectExchangeABI, nil, nil, nil)

	niftyConnectEvents := make([]interface{}, 0, len(logs))
	for _, log := range logs {
		switch log.Topics[0].String() {
		case adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartOneEvent].ID.String():
			var orderApprovedPartOne OrderApprovedPartOne
			err := niftyConnectContract.UnpackLog(&orderApprovedPartOne, orderApprovedPartOneEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", orderApprovedPartOneEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, orderApprovedPartOne.toModelType(adaptorInst.blockchain, log.TxHash.String(), int64(log.BlockNumber), int64(log.Index)))
		case adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartTwoEvent].ID.String():
			var orderApprovedPartTwo OrderApprovedPartTwo
			err := niftyConnectContract.UnpackLog(&orderApprovedPartTwo, orderApprovedPartTwoEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", orderApprovedPartTwoEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, orderApprovedPartTwo.toModelType(adaptorInst.blockchain, log.TxHash.String(), int64(log.BlockNumber), int64(log.Index)))
		case adaptorInst.niftyConnectExchangeABI.Events[orderCancelledEvent].ID.String():
			var orderCancelled OrderCancelled
			err := niftyConnectContract.UnpackLog(&orderCancelled, orderCancelledEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", orderCancelledEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, orderCancelled.toModelType(adaptorInst.blockchain, log.TxHash.String(), int64(log.BlockNumber), int64(log.Index)))
		case adaptorInst.niftyConnectExchangeABI.Events[ordersMatchedEvent].ID.String():
			var ordersMatched OrdersMatched
			err := niftyConnectContract.UnpackLog(&ordersMatched, ordersMatchedEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", ordersMatchedEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, ordersMatched.toModelType(adaptorInst.blockchain, log.TxHash.String(), int64(log.BlockNumber), int64(log.Index)))
		case adaptorInst.niftyConnectExchangeABI.Events[nonceIncrementedEvent].ID.String():
			var nonceIncremented NonceIncremented
			err := niftyConnectContract.UnpackLog(&nonceIncremented, nonceIncrementedEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", nonceIncrementedEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, nonceIncremented.toModelType(adaptorInst.blockchain, log.TxHash.String(), int64(log.BlockNumber), int64(log.Index)))
		default:
			localcmm.Logger.Errorf("Unexpected event topic: %s", log.Topics[0].String())
		}
	}
	return niftyConnectEvents, nil
}
