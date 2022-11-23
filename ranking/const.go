package ranking

const (
	nftScanApi        = "https://restapi.nftscan.com/api/v2/nifty/orders?offset=%d&order_stat=Finalized&limit=%d"
	nftScanApiAll     = "https://restapi.nftscan.com/api/v2/nifty/orders?offset=%d&limit=%d"
	xApiKey           = "wU2sELIJ"
	perPage           = 100
	retryInterval     = 10
	twitterPointsFile = "twitterPoints.cfg"
)

var (
	allScanMakerOrders   = make(map[string]OrderInfo)
	allScanTakerOrders   = make(map[string]OrderInfo)
	allUserOrders        = make(map[string]map[string]map[string]interface{})
	userPoints           = make(map[string]int)
	userPointsSort       = make(map[string]int)
	allUserListingOrders = make(map[string]map[string]map[string]interface{})
	userListingPoints    = make(map[string]map[string]map[string]interface{})
)
