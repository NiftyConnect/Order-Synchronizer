package synchronizer

import "github.com/ethereum/go-ethereum/common"

const (
	observerPruneInterval = 60
	contextTimeout        = 30

	pruneHeightGap = 1000
)

var (
	zeroAddress = common.Address{}
)
