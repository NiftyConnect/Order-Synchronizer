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

func NewClientMgr(cfg *config.ChainDetail, blockchain string) (*ClientMgr, error) {
	clientMgr := &ClientMgr{
		idx:          0,
		clientSetLen: int64(len(cfg.Providers)),
		providers:    cfg.Providers,
		blockchain:   blockchain,
		clientSet:    make([]*ethclient.Client, 0, len(cfg.Providers)),
	}
	for _, provider := range cfg.Providers {
		ethClient, err := ethclient.Dial(provider)
		if err != nil {
			return nil, err
		}
		clientMgr.clientSet = append(clientMgr.clientSet, ethClient)
	}
	return clientMgr, nil
}

func (clientMgr *ClientMgr) GetClient() *ethclient.Client {
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
