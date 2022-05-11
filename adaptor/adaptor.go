package adaptor
import (
	"context"
	"fmt"
	"github.com/niftyConnect/order-synchronizer/alert"
	"math/big"
	"strings"
	"time"
	
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	
	"github.com/niftyConnect/order-synchronizer/database"
	localcmm "github.com/niftyConnect/order-synchronizer/common"
	"github.com/niftyConnect/order-synchronizer/config"
	"github.com/niftyConnect/order-synchronizer/adaptor/contracts"
)

var (
	zeroAddress        = common.Address{}
)

type Adaptor struct {
	blockchain string
	cfg        *config.Config

	niftyConnectExchangeABI abi.ABI

	clientHub ClientHub
}

func NewExecutorV0(cfg *config.Config, blockchain string, clientHub ClientHub) (*Adaptor, error) {
	niftyConnectExchangeABIV, err := abi.JSON(strings.NewReader(contracts.NiftyconnectexchangeABI))
	if err != nil {
		return nil, fmt.Errorf("marshal abi error")
	}

	return &Adaptor{
		blockchain:              blockchain,
		cfg:                     cfg,

		clientHub:               clientHub,
		niftyConnectExchangeABI: niftyConnectExchangeABIV,
	}, nil
}

func (adaptorInst *Adaptor) GetHeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return adaptorInst.clientHub.GetClient(adaptorInst.blockchain).HeaderByNumber(ctx, number)
}

func (adaptorInst *Adaptor) GetLogs(header *types.Header, contractAddrs []common.Address) ([]database.NiftyConnectOrder, error) {
	if len(contractAddrs) == 0 {
		return nil, nil
	}
	return adaptorInst.getERC721TransferEvent(header, contractAddrs)
}

func (adaptorInst *Adaptor) getERC721TransferEvent(header *types.Header, contractAddrs []common.Address) ([]database.NiftyConnectOrder, error) {
	topics := [][]common.Hash{{adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartOneEvent].ID}}

	blockHash := header.Hash()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Second)
	defer cancel()

	logs, err := adaptorInst.clientHub.GetClient(adaptorInst.blockchain).FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: contractAddrs,
	})
	if err != nil {
		return nil, err
	}

	eventModels := make([]database.NiftyConnectOrder, 0, len(logs))
	for _, log := range logs {
		erc721Contract := bind.NewBoundContract(zeroAddress, adaptorInst.niftyConnectExchangeABI, nil, nil, nil)
		var erc721TransferEvent ERC721Transfer
		err = erc721Contract.UnpackLog(&erc721TransferEvent, orderApprovedPartOneEvent, log)
		if err != nil {
			alert.SendTelegramMessage(adaptorInst.cfg, fmt.Sprintf("parse event parameters with nftABIV1 err: %s, txHash %s", err.Error(), log.TxHash.String()))
			localcmm.Logger.Errorf("parse event parameters with nftABIV1 err: %s, txHash %s", err.Error(), log.TxHash.String())
		}
		localcmm.Logger.Debugf("blockchain %s, ERC721 Transfer, from: %s, to: %s, tokenId: %s， tokenURI: %s, txHash %s",
			adaptorInst.blockchain, erc721TransferEvent.From.String(), erc721TransferEvent.To.String(), erc721TransferEvent.TokenId.Text(16), tokenURI, log.TxHash.String())
		eventModel := database.NiftyConnectOrder{
			EventId:          generateEventId(log.BlockNumber, uint64(log.Index), 0),
			NFTAddress:       log.Address.String(),
			TokenId:          bigIntToHexString(erc721TransferEvent.TokenId),
			TokenURI:         tokenURI,
			Amount:           erc721TransferAmount,
			SrcAddress:       erc721TransferEvent.From.String(),
			DstAddress:       erc721TransferEvent.To.String(),
			TxHash:           log.TxHash.String(),
			Height:           int64(log.BlockNumber),
			NFTTokenStandard: int64(config.ERC721Standard),
			Blockchain:       adaptorInst.blockchain,
			CreateTime:       time.Now().Unix(),
			UpdateTime:       time.Now().Unix(),
		}
		eventModels = append(eventModels, eventModel)
	}
	return eventModels, nil
}

func (adaptorInst *Adaptor) ExtractNFTTransferEvent(fromHeight, toHeight int64, contractAddr common.Address, includeTokenURI bool) ([]database.NiftyConnectOrder, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Second)
	defer cancel()

	topics := [][]common.Hash{{adaptorInst.niftyConnectExchangeABI.Events[orderApprovedPartOneEvent].ID}}
	logs, err := adaptorInst.clientHub.GetClient(adaptorInst.blockchain).FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		FromBlock: big.NewInt(fromHeight),
		ToBlock:   big.NewInt(toHeight),
		Topics:    topics,
		Addresses: []common.Address{contractAddr},
	})
	if err != nil {
		return nil, err
	}

	eventModels := make([]database.NiftyConnectOrder, 0, len(logs))
	for _, log := range logs {
		erc721Contract := bind.NewBoundContract(zeroAddress, adaptorInst.nftABIV1, nil, nil, nil)
		var erc721TransferEvent ERC721Transfer
		err = erc721Contract.UnpackLog(&erc721TransferEvent, orderApprovedPartOneEvent, log)
		if err != nil {
			alert.SendTelegramMessage(adaptorInst.cfg, fmt.Sprintf("parse event parameters with nftABIV1 err: %s, txHash %s", err.Error(), log.TxHash.String()))
			localcmm.Logger.Errorf("parse event parameters with nftABIV1 err: %s, txHash %s", err.Error(), log.TxHash.String())

			erc721Contract = bind.NewBoundContract(zeroAddress, adaptorInst.nftABIV2, nil, nil, nil)
			err = erc721Contract.UnpackLog(&erc721TransferEvent, orderApprovedPartOneEvent, log)
			if err != nil {
				alert.SendTelegramMessage(adaptorInst.cfg, fmt.Sprintf("parse event parameters with nftABIV2 err: %s, txHash %s", err.Error(), log.TxHash.String()))
				return nil, fmt.Errorf("parse event parameters with nftABIV2 err: %s, txHash %s", err.Error(), log.TxHash.String())
			}
		}
		localcmm.Logger.Debugf("blockchain %s, ERC721 Transfer, from: %s, to: %s, tokenId: %s， tokenURI: %s, txHash %s",
			adaptorInst.blockchain, erc721TransferEvent.From.String(), erc721TransferEvent.To.String(), erc721TransferEvent.TokenId.Text(16), tokenURI, log.TxHash.String())
		eventModel := database.NiftyConnectOrder{
			EventId:          generateEventId(log.BlockNumber, uint64(log.Index), 0),
			NFTAddress:       log.Address.String(),
			TokenId:          bigIntToHexString(erc721TransferEvent.TokenId),
			TokenURI:         tokenURI,
			Amount:           erc721TransferAmount,
			SrcAddress:       erc721TransferEvent.From.String(),
			DstAddress:       erc721TransferEvent.To.String(),
			TxHash:           log.TxHash.String(),
			Height:           int64(log.BlockNumber),
			NFTTokenStandard: int64(config.ERC721Standard),
			Blockchain:       adaptorInst.blockchain,
			CreateTime:       time.Now().Unix(),
			UpdateTime:       time.Now().Unix(),
		}
		eventModels = append(eventModels, eventModel)
	}
	return eventModels, nil
}

func (adaptorInst *Adaptor) ExtractEvents(initHeightStep int64, latestHeight, toHeight int64, contractAddress common.Address) ([]database.NiftyConnectOrder, error) {
	heightStep := initHeightStep
	startHeight := latestHeight + 1
	endHeight := startHeight + heightStep - 1
	result := make([]database.NiftyConnectOrder, 0)

	localcmm.Logger.Infof("Blockchain %s, nftAddress %s, delta filter event from %d to %d, initHeightStep %d",
		adaptorInst.blockchain, contractAddress.String(), latestHeight, toHeight, initHeightStep)
	for startHeight <= toHeight {
		if endHeight > toHeight {
			endHeight = toHeight
		}
		localcmm.Logger.Debugf("Blockchain %s, nftAddress %s, filter event between %d, %d, height step %d",
			adaptorInst.blockchain, contractAddress.String(), startHeight, endHeight, heightStep)
	retry:
		events, err := func() (transferEvents []database.NiftyConnectOrder, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf(fmt.Sprintf("%v", r))
				}
			}()
			transferEvents, err = adaptorInst.ExtractNFTTransferEvent(startHeight, endHeight, contractAddress, true)
			if err != nil {
				return nil, fmt.Errorf("blockchain %s, nft_address %s, extract transfer event err: %s", adaptorInst.blockchain, contractAddress.String(), err)
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
	localcmm.Logger.Debugf("Complete syncing recent events for %s", contractAddress.String())
	return result, nil
}
