package synchronizer

import "github.com/ethereum/go-ethereum/common"

const (
	pruneInterval  = 60
	pruneHeightGap = 1000

	expireInterval = 60
)

var (
	zeroAddress = common.Address{}
)
