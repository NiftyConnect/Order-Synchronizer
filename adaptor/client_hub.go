package adaptor

import (
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"

	localcmm "github.com/niftyConnect/order-synchronizer/common"
	"github.com/niftyConnect/order-synchronizer/config"
)

type ClientMgr struct {
	mutex        sync.RWMutex
	idx          int64
	clientSetLen int64
	blockchain   string
	providers    []string
	clientSet    []*ethclient.Client
}

type ClientHub map[string]*ClientMgr

func NewClientHub(cfg *config.Config) (ClientHub, error) {
	clientHub := make(map[string]*ClientMgr)
	for blockchain, chainDetail := range cfg.ChainConfig.ChainDetails {
		clientMgr := &ClientMgr{
			idx:          0,
			clientSetLen: int64(len(chainDetail.Providers)),
			providers:    chainDetail.Providers,
			blockchain:   blockchain,
			clientSet:    make([]*ethclient.Client, 0, len(chainDetail.Providers)),
		}
		for _, provider := range chainDetail.Providers {
			ethClient, err := ethclient.Dial(provider)
			if err != nil {
				return nil, err
			}
			clientMgr.clientSet = append(clientMgr.clientSet, ethClient)
		}
		clientHub[blockchain] = clientMgr
	}
	return clientHub, nil
}

func (clientHub ClientHub) GetClient(blockchain string) *ethclient.Client {
	clientMgr := clientHub[blockchain]
	clientMgr.mutex.Lock()
	client := clientMgr.clientSet[clientMgr.idx]
	localcmm.Logger.Debugf("blockchain %s, client provider: %s", clientMgr.blockchain, clientMgr.providers[clientMgr.idx])
	clientMgr.idx++
	if clientMgr.idx >= clientMgr.clientSetLen {
		clientMgr.idx = 0
	}
	clientMgr.mutex.Unlock()
	return client
}
