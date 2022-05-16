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
	cfg                     *config.Config
	niftyConnectExchangeABI abi.ABI
	niftyConnectAddress     common.Address
	clientHub               ClientHub
}

func NewAdaptor(cfg *config.Config, blockchain string, clientHub ClientHub) (*Adaptor, error) {
	niftyConnectExchangeABIV, err := abi.JSON(strings.NewReader(contracts.NiftyconnectexchangeABI))
	if err != nil {
		return nil, fmt.Errorf("marshal abi error")
	}
	chainDetails, ok := cfg.ChainConfig.ChainDetails[blockchain]
	if !ok {
		return nil, fmt.Errorf("missing chain details for %s", blockchain)
	}
	return &Adaptor{
		blockchain:              blockchain,
		cfg:                     cfg,
		clientHub:               clientHub,
		niftyConnectExchangeABI: niftyConnectExchangeABIV,
		niftyConnectAddress:     chainDetails.ExchangeContract,
	}, nil
}

func (adaptorInst *Adaptor) GetHeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return adaptorInst.clientHub.GetClient(adaptorInst.blockchain).HeaderByNumber(ctx, number)
}

func (adaptorInst *Adaptor) GetLogs(header *types.Header) ([]interface{}, error) {
	return adaptorInst.getNiftyConnectEventAtBlock(header)
}

func (adaptorInst *Adaptor) getNiftyConnectEventAtBlock(header *types.Header) ([]interface{}, error) {
	topics := [][]common.Hash{{
		adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartOneEvent].ID(),
		adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartTwoEvent].ID(),
		adaptorInst.niftyConnectExchangeABI.Events[orderCancelledEvent].ID(),
		adaptorInst.niftyConnectExchangeABI.Events[ordersMatchedEvent].ID(),
		adaptorInst.niftyConnectExchangeABI.Events[nonceIncrementedEvent].ID(),
	}}

	blockHash := header.Hash()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Second)
	defer cancel()

	logs, err := adaptorInst.clientHub.GetClient(adaptorInst.blockchain).FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []common.Address{adaptorInst.niftyConnectAddress},
	})
	if err != nil {
		return nil, err
	}

	niftyConnectContract := bind.NewBoundContract(zeroAddress, adaptorInst.niftyConnectExchangeABI, nil, nil, nil)

	niftyConnectEvents := make([]interface{}, 0, len(logs))
	for _, log := range logs {
		switch log.Topics[0].String() {
		case adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartOneEvent].ID().String():
			var orderApprovedPartOne OrderApprovedPartOne
			err = niftyConnectContract.UnpackLog(&orderApprovedPartOne, orderApprovedPartOneEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", orderApprovedPartOneEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, orderApprovedPartOne)
		case adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartTwoEvent].ID().String():
			var orderApprovedPartTwo OrderApprovedPartTwo
			err = niftyConnectContract.UnpackLog(&orderApprovedPartTwo, orderApprovedPartTwoEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", orderApprovedPartTwoEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, orderApprovedPartTwo)
		case adaptorInst.niftyConnectExchangeABI.Events[orderCancelledEvent].ID().String():
			var orderCancelled OrderCancelled
			err = niftyConnectContract.UnpackLog(&orderCancelled, orderCancelledEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", orderCancelledEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, orderCancelled)
		case adaptorInst.niftyConnectExchangeABI.Events[ordersMatchedEvent].ID().String():
			var ordersMatched OrdersMatched
			err = niftyConnectContract.UnpackLog(&ordersMatched, ordersMatchedEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", ordersMatchedEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, ordersMatched)
		case adaptorInst.niftyConnectExchangeABI.Events[nonceIncrementedEvent].ID().String():
			var nonceIncremented NonceIncremented
			err = niftyConnectContract.UnpackLog(&nonceIncremented, nonceIncrementedEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", nonceIncrementedEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, nonceIncremented)
		default:
			localcmm.Logger.Errorf("Unexpected event topic: %s", log.Topics[0].String())
		}
	}
	return niftyConnectEvents, nil
}

func (adaptorInst *Adaptor) getNiftyConnectEventFromBlockRange(fromHeight, toHeight int64) ([]interface{}, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Second)
	defer cancel()

	topics := [][]common.Hash{{
		adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartOneEvent].ID(),
		adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartTwoEvent].ID(),
		adaptorInst.niftyConnectExchangeABI.Events[orderCancelledEvent].ID(),
		adaptorInst.niftyConnectExchangeABI.Events[ordersMatchedEvent].ID(),
		adaptorInst.niftyConnectExchangeABI.Events[nonceIncrementedEvent].ID(),
	}}

	logs, err := adaptorInst.clientHub.GetClient(adaptorInst.blockchain).FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		FromBlock: big.NewInt(fromHeight),
		ToBlock:   big.NewInt(toHeight),
		Topics:    topics,
		Addresses: []common.Address{adaptorInst.niftyConnectAddress},
	})
	if err != nil {
		return nil, err
	}

	niftyConnectContract := bind.NewBoundContract(zeroAddress, adaptorInst.niftyConnectExchangeABI, nil, nil, nil)

	niftyConnectEvents := make([]interface{}, 0, len(logs))
	for _, log := range logs {
		switch log.Topics[0].String() {
		case adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartOneEvent].ID().String():
			var orderApprovedPartOne OrderApprovedPartOne
			err = niftyConnectContract.UnpackLog(&orderApprovedPartOne, orderApprovedPartOneEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", orderApprovedPartOneEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, orderApprovedPartOne)
		case adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartTwoEvent].ID().String():
			var orderApprovedPartTwo OrderApprovedPartTwo
			err = niftyConnectContract.UnpackLog(&orderApprovedPartTwo, orderApprovedPartTwoEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", orderApprovedPartTwoEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, orderApprovedPartTwo)
		case adaptorInst.niftyConnectExchangeABI.Events[orderCancelledEvent].ID().String():
			var orderCancelled OrderCancelled
			err = niftyConnectContract.UnpackLog(&orderCancelled, orderCancelledEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", orderCancelledEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, orderCancelled)
		case adaptorInst.niftyConnectExchangeABI.Events[ordersMatchedEvent].ID().String():
			var ordersMatched OrdersMatched
			err = niftyConnectContract.UnpackLog(&ordersMatched, ordersMatchedEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", ordersMatchedEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, ordersMatched)
		case adaptorInst.niftyConnectExchangeABI.Events[nonceIncrementedEvent].ID().String():
			var nonceIncremented NonceIncremented
			err = niftyConnectContract.UnpackLog(&nonceIncremented, nonceIncrementedEvent, log)
			if err != nil {
				localcmm.Logger.Errorf("UnpackLog event %s err: %s", nonceIncrementedEvent, err.Error())
				continue
			}
			niftyConnectEvents = append(niftyConnectEvents, nonceIncremented)
		default:
			localcmm.Logger.Errorf("Unexpected event topic: %s", log.Topics[0].String())
		}
	}
	return niftyConnectEvents, nil
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
		localcmm.Logger.Debugf("Blockchain %s, filter event between %d, %d, height step %d",
			adaptorInst.blockchain, startHeight, endHeight, heightStep)
	retry:
		events, err := func() (transferEvents []interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf(fmt.Sprintf("%v", r))
				}
			}()
			transferEvents, err = adaptorInst.getNiftyConnectEventFromBlockRange(startHeight, endHeight)
			if err != nil {
				return nil, fmt.Errorf("blockchain %s, extract NiftyConnect event err: %s", adaptorInst.blockchain, err)
			}
			return transferEvents, nil
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
	localcmm.Logger.Debugf("Complete syncing recent events")
	return result, nil
}
