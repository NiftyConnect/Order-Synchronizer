package ranking

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/niftyConnect/order-synchronizer/config"
)

type NftScanResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data []OrderInfo `json:"data"`
}
type OrderInfo struct {
	Order_hash                  string `json:"order_hash"`
	Tx_hash                     string `json:"tx_hash"`
	Tx_time                     int    `json:"tx_time"`
	Exchange                    string `json:"exchange"`
	Maker                       string `json:"maker"`
	Taker                       string `json:"taker"`
	Maker_relayer_fee_recipient string `json:"maker_relayer_fee_recipient"`
	Taker_relayer_fee_recipient string `json:"taker_relayer_fee_recipient"`
	Sale_side                   string `json:"sale_side"`
	Sale_kind                   string `json:"sale_kind"`
	Nft_address                 string `json:"nft_address"`
	Token_id                    string `json:"token_id"`
	Ipfs_hash                   string `json:"ipfs_hash"`
	Calldata                    string `json:"calldata"`
	Replacement_pattern         string `json:"replacement_pattern"`
	Static_target               string `json:"static_target"`
	Static_extradata            string `json:"static_extradata"`
	Payment_token               string `json:"payment_token"`
	Base_price                  string `json:"base_price"`
	Extra                       string `json:"extra"`
	Listing_time                int    `json:"listing_time"`
	Expiration_time             int    `json:"expiration_time"`
	Salt                        string `json:"salt"`
	Trade_price                 string `json:"trade_price"`
	Rss_metadata                string `json:"rss_metadata"`
	Order_stat                  string `json:"order_stat"`
	Order_change_hash           string `json:"order_change_hash"`
	Order_change_time           int    `json:"order_change_time"`
}

type UserOrderInfo struct {
	Tx_hash     string
	Side        string //"buy", "sell"
	Nft_address string
	Token_id    string
	Maker       string
	Taker       string
	Tx_time     int
}

func Listing(cfg *config.Config) {
	fmt.Println("================= General NFT calculating =================")
	for {
		offset := 0
		clearListingMaps()
		for {
			scanOrders, err := readFromNftScan(offset, true)

			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("retry in ", retryInterval, "seconds")
				time.Sleep(time.Duration(retryInterval) * time.Second)
				continue
			}
			ordersNum := len(scanOrders)

			for _, order := range scanOrders {

				if nil == allUserListingOrders[order.Maker] {
					allUserListingOrders[order.Maker] = make(map[string]map[string]interface{})
				}
				if orderOngoing(order, cfg) {

					ordersKey := ""
					if order.Order_stat == "Finalized" {
						ordersKey = order.Tx_hash
						oldKey := order.Maker + order.Nft_address + order.Token_id
						if len(allUserListingOrders[order.Maker][oldKey]) != 0 {
							continue
						}
					} else {
						ordersKey = order.Maker + order.Nft_address + order.Token_id
					}
					if len(allUserListingOrders[order.Maker][ordersKey]) == 0 {

						if len(allUserListingOrders[order.Maker][order.Tx_hash]) != 0 {
							continue
						}
						allUserListingOrders[order.Maker][ordersKey] = make(map[string]interface{})
						allUserListingOrders[order.Maker][ordersKey]["Side"] = "sell"
						allUserListingOrders[order.Maker][ordersKey]["Maker"] = order.Maker
						allUserListingOrders[order.Maker][ordersKey]["Taker"] = order.Taker
						allUserListingOrders[order.Maker][ordersKey]["Tx_time"] = order.Tx_time
						allUserListingOrders[order.Maker][ordersKey]["Tx_hash"] = order.Tx_hash
						allUserListingOrders[order.Maker][ordersKey]["Token_id"] = order.Token_id
						allUserListingOrders[order.Maker][ordersKey]["Nft_address"] = order.Nft_address
						allUserListingOrders[order.Maker][ordersKey]["Listing_time"] = order.Listing_time
						allUserListingOrders[order.Maker][ordersKey]["Order_change_time"] = order.Order_change_time
						allUserListingOrders[order.Maker][ordersKey]["Order_stat"] = order.Order_stat

					}
				}
			}
			if ordersNum < perPage {
				break
			}
			offset += perPage
		} //for
		userListingPoints = make(map[string]map[string]map[string]interface{})
		for user, userOrders := range allUserListingOrders {
			for _, order := range userOrders {
				if nil == userListingPoints[user] {
					userListingPoints[user] = make(map[string]map[string]interface{})
				}
				if nil == userListingPoints[user]["cb"] {
					userListingPoints[user]["cb"] = make(map[string]interface{})
					userListingPoints[user]["cb"]["bc"] = 0
					userListingPoints[user]["cb"]["nm"] = 0
				}
				if nil == userListingPoints[user]["ob"] {
					userListingPoints[user]["ob"] = make(map[string]interface{})
					userListingPoints[user]["ob"]["bc"] = 0
					userListingPoints[user]["ob"]["nm"] = 0
				}
				if nil == userListingPoints[user] {
					userListingPoints[user] = make(map[string]map[string]interface{})
					userListingPoints[user]["cb"] = make(map[string]interface{})
					userListingPoints[user]["cb"]["bc"] = 0
					userListingPoints[user]["cb"]["nm"] = 0
					userListingPoints[user]["ob"] = make(map[string]interface{})
					userListingPoints[user]["ob"]["bc"] = 0
					userListingPoints[user]["ob"]["nm"] = 0
				}
				listingTime, _ := strconv.Atoi(fmtStrFromInterface(order["Listing_time"]))
				if whiteListed(user, cfg.RankingConfig.WhiteList) {
					//白名单用户
					if listingTime > cfg.RankingConfig.CbStart && listingTime <= cfg.RankingConfig.CbEnd {
						if isBlueChip(fmtStrFromInterface(order["Nft_address"]), cfg.RankingConfig.BlueChip) {
							if userListingPoints[user]["cb"]["bc"] != 0 {
								p := fmtStrFromInterface(userListingPoints[user]["cb"]["bc"])
								pNum, _ := strconv.Atoi(p)
								userListingPoints[user]["cb"]["bc"] = pNum + 1
							} else {
								userListingPoints[user]["cb"]["bc"] = 1
							}
						} else {
							if userListingPoints[user]["cb"]["nm"] != 0 {
								p := fmtStrFromInterface(userListingPoints[user]["cb"]["nm"])
								pNum, _ := strconv.Atoi(p)
								userListingPoints[user]["cb"]["nm"] = pNum + 1
							} else {
								userListingPoints[user]["cb"]["nm"] = 1
							}
						}
					} else if listingTime > cfg.RankingConfig.ObStart && listingTime <= cfg.RankingConfig.ObEnd {
						if isBlueChip(fmtStrFromInterface(order["Nft_address"]), cfg.RankingConfig.BlueChip) {
							if userListingPoints[user]["ob"]["bc"] != 0 {
								p := fmtStrFromInterface(userListingPoints[user]["ob"]["bc"])
								pNum, _ := strconv.Atoi(p)
								userListingPoints[user]["ob"]["bc"] = pNum + 1
							} else {
								userListingPoints[user]["ob"]["bc"] = 1
							}
						} else {
							if userListingPoints[user]["ob"]["nm"] != 0 {
								p := fmtStrFromInterface(userListingPoints[user]["ob"]["nm"])
								pNum, _ := strconv.Atoi(p)
								userListingPoints[user]["ob"]["nm"] = pNum + 1
							} else {
								userListingPoints[user]["ob"]["nm"] = 1
							}
						}
					}
				} else {
					//普通用户
					if listingTime > cfg.RankingConfig.ObStart && listingTime <= cfg.RankingConfig.ObEnd {
						if isBlueChip(fmtStrFromInterface(order["Nft_address"]), cfg.RankingConfig.BlueChip) {
							if userListingPoints[user]["ob"]["bc"] != 0 {
								p := fmtStrFromInterface(userListingPoints[user]["ob"]["bc"])
								pNum, _ := strconv.Atoi(p)
								userListingPoints[user]["ob"]["bc"] = pNum + 1
							} else {
								userListingPoints[user]["ob"]["bc"] = 1
							}
						} else {
							if userListingPoints[user]["ob"]["nm"] != 0 {
								p := fmtStrFromInterface(userListingPoints[user]["ob"]["nm"])
								pNum, _ := strconv.Atoi(p)
								userListingPoints[user]["ob"]["nm"] = pNum + 1
							} else {
								userListingPoints[user]["ob"]["nm"] = 1
							}
						}
					}
				}
			}
		} //for
		userNfts := make(map[string]int, len(userListingPoints))
		for user, listingPoints := range userListingPoints {
			nftNum := 0
			if whiteListed(user, cfg.RankingConfig.WhiteList) {
				if listingPoints["cb"]["bc"] != 0 {
					bc := fmtStrFromInterface(listingPoints["cb"]["bc"])
					bcNum, _ := strconv.Atoi(bc)
					nftNum += 4 * bcNum
				}
				if listingPoints["cb"]["nm"] != 0 {
					nm := fmtStrFromInterface(listingPoints["cb"]["nm"])
					nmNum, _ := strconv.Atoi(nm)
					nftNum += 2 * nmNum
				}
				if listingPoints["ob"]["bc"] != 0 {
					bc := fmtStrFromInterface(listingPoints["ob"]["bc"])
					bcNum, _ := strconv.Atoi(bc)
					nftNum += 3 * bcNum
				}
				if listingPoints["ob"]["nm"] != 0 {
					nm := fmtStrFromInterface(listingPoints["ob"]["nm"])
					nmNum, _ := strconv.Atoi(nm)
					nftNum += nmNum
				}
			} else {
				if listingPoints["ob"]["bc"] != 0 {
					bc := fmtStrFromInterface(listingPoints["ob"]["bc"])
					bcNum, _ := strconv.Atoi(bc)
					nftNum += 3 * bcNum
				}
				if listingPoints["ob"]["nm"] != 0 {
					nm := fmtStrFromInterface(listingPoints["ob"]["nm"])
					nmNum, _ := strconv.Atoi(nm)
					nftNum += nmNum
				}
			}
			userNfts[user] = nftNum
		} //for
		iUsers := 0
		for key := range userNfts {
			if userNfts[key] > 0 {
				iUsers++
			}
		}
		keys := make([]string, iUsers)
		for key := range userNfts {
			if userNfts[key] > 0 {
				keys = append(keys, key)
			}
		}
		sort.SliceStable(keys, func(i, j int) bool {
			return userNfts[keys[i]] > userPoints[keys[j]]
		})
		userNftsArray := make([]map[string]interface{}, iUsers)
		indexArray := 0
		for _, k := range keys {
			if userNfts[k] > 0 {
				userNftsArray[indexArray] = make(map[string]interface{})
				userNftsArray[indexArray]["key"] = k
				userNftsArray[indexArray]["value"] = userNfts[k]
				indexArray++
			}
		}
		str, _ := json.Marshal(userNftsArray)
		err := ioutil.WriteFile(cfg.RankingConfig.NftFile, str, 0644)
		if nil != err {
			fmt.Println("failed to write general nft data file ", err, cfg.RankingConfig.NftFile)
		}
		time.Sleep(time.Duration(cfg.RankingConfig.RankingInterval) * time.Second)
	}
}

func orderOngoing(order OrderInfo, cfg *config.Config) bool {
	if order.Listing_time < cfg.RankingConfig.CbStart {
		return false
	}
	if order.Order_stat == "Finalized" {
		return true
	} else if order.Order_stat == "Canceled" {
		return true
	} else {

		now := int(time.Now().Unix())
		if order.Expiration_time < now {
			return (order.Expiration_time-order.Listing_time > cfg.RankingConfig.MinListedTime)
		} else {
			return (now-order.Listing_time > cfg.RankingConfig.MinListedTime)
		}

	}
}

func Start(cfg *config.Config) {

	//隔一段时间重新跑
	for {
		clearMaps()
		t1 := time.Now().Unix()

		allUserOrders = readAllOrders()
		userPoints = calcPoints(allUserOrders, cfg)

		//// >>>>>>>>>>>>>>>>>>>>> 计算交易积分排行榜，读取twitter 20 分 >>>>>>>>>>>>>>>>>>>>>
		twitterFileContent, err := ioutil.ReadFile(twitterPointsFile)
		var twitterUsers []string
		if err != nil {
			fmt.Println("error reading twitterConfig file", err)
		} else {
			if len(twitterFileContent) > 0 {
				twitterUsers = strings.Split(string(twitterFileContent), "\n")
				for _, v := range twitterUsers {
					userPoints[v] += 20
				}
			}
		}
		//// <<<<<<<<<<<<<<<<<<<<< 计算交易积分排行榜，读取twitter 20 分 <<<<<<<<<<<<<<<<<<<<<

		// >>>>>>>>>>>>>>>>>>>>> 排序交易积分排行榜 >>>>>>>>>>>>>>>>>>>>>
		keys := make([]string, 0, len(userPoints))
		for key := range userPoints {
			keys = append(keys, key)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			return userPoints[keys[i]] > userPoints[keys[j]]
		})
		// <<<<<<<<<<<<<<<<<<<<< 排序交易积分排行榜 <<<<<<<<<<<<<<<<<<<<<

		//===================== 交易排行榜写入文件给rest读 =====================
		writeUserTradePoints(userPoints, keys, cfg)

		// >>>>>>>>>>>>>>>>>>>>> 开始 交易量排行榜 >>>>>>>>>>>>>>>>>>>>>
		userVolumes := calcVolume(allUserOrders, cfg)
		// >>>>>>>>>>>>>>>>>>>>> 排序交易量排行榜 >>>>>>>>>>>>>>>>>>>>>
		/* 		keys = make([]string, 0, len(userVolumes))
		   		for key := range userVolumes {
		   			keys = append(keys, key)
		   		} */
		keys = make([]string, 0, len(userPoints))
		for key := range userPoints {
			keys = append(keys, key)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			return userPoints[keys[i]] > userPoints[keys[j]]
		})
		writeUserVolumes(userVolumes, userPoints, keys, cfg)
		// <<<<<<<<<<<<<<<<<<<<< 排序交易量排行榜 <<<<<<<<<<<<<<<<<<<<<

		// <<<<<<<<<<<<<<<<<<<<< 结束 交易量排行榜 <<<<<<<<<<<<<<<<<<<<<
		t2 := time.Now().Unix()
		fmt.Println("It took", (t2 - t1), "seconds to make ranking")
		time.Sleep(time.Duration(cfg.RankingConfig.RankingInterval) * time.Second)
	}
}

func readAllOrders() map[string]map[string]map[string]interface{} {
	offset := 0
	_allUserOrders := make(map[string]map[string]map[string]interface{})
	// >>>>>>>>>>>>>>>>>>>>> 开始 从 nftscan 读取全部成交定单 >>>>>>>>>>>>>>>>>>>>>
	for {
		scanOrders, err := readFromNftScan(offset, false)

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("retry in ", retryInterval, "seconds")
			time.Sleep(time.Duration(retryInterval) * time.Second)
			continue
		}
		ordersNum := len(scanOrders)

		for _, order := range scanOrders {

			if nil == _allUserOrders[order.Maker] {
				//fmt.Println("add maker map")
				_allUserOrders[order.Maker] = make(map[string]map[string]interface{})
			}
			//fmt.Println(allUserOrders[order.Maker][order.Tx_hash])
			if len(_allUserOrders[order.Maker][order.Tx_hash]) == 0 {
				//fmt.Println("add maker hash => order")
				_allUserOrders[order.Maker][order.Tx_hash] = make(map[string]interface{})
				_allUserOrders[order.Maker][order.Tx_hash]["Side"] = "sell"
				_allUserOrders[order.Maker][order.Tx_hash]["Maker"] = order.Maker
				_allUserOrders[order.Maker][order.Tx_hash]["Taker"] = order.Taker
				_allUserOrders[order.Maker][order.Tx_hash]["Tx_time"] = order.Tx_time
				_allUserOrders[order.Maker][order.Tx_hash]["Tx_hash"] = order.Tx_hash
				_allUserOrders[order.Maker][order.Tx_hash]["Token_id"] = order.Token_id
				_allUserOrders[order.Maker][order.Tx_hash]["Nft_address"] = order.Nft_address
				_allUserOrders[order.Maker][order.Tx_hash]["Listing_time"] = order.Listing_time
				_allUserOrders[order.Maker][order.Tx_hash]["Order_change_time"] = order.Order_change_time
				_allUserOrders[order.Maker][order.Tx_hash]["Trade_price"] = order.Trade_price
				//fmt.Println("scan sell order change time:", order.Order_change_time)

			}
			if nil == _allUserOrders[order.Taker] {
				//fmt.Println("add taker map")
				_allUserOrders[order.Taker] = make(map[string]map[string]interface{})
			}

			if len(_allUserOrders[order.Taker][order.Tx_hash]) == 0 {
				//fmt.Println("add taker hask => order")
				_allUserOrders[order.Taker][order.Tx_hash] = make(map[string]interface{})
				_allUserOrders[order.Taker][order.Tx_hash]["Side"] = "buy"
				_allUserOrders[order.Taker][order.Tx_hash]["Maker"] = order.Maker
				_allUserOrders[order.Taker][order.Tx_hash]["Taker"] = order.Taker
				_allUserOrders[order.Taker][order.Tx_hash]["Tx_time"] = order.Tx_time
				_allUserOrders[order.Taker][order.Tx_hash]["Tx_hash"] = order.Tx_hash
				_allUserOrders[order.Taker][order.Tx_hash]["Token_id"] = order.Token_id
				_allUserOrders[order.Taker][order.Tx_hash]["Nft_address"] = order.Nft_address
				_allUserOrders[order.Taker][order.Tx_hash]["Listing_time"] = order.Listing_time
				_allUserOrders[order.Taker][order.Tx_hash]["Order_change_time"] = order.Order_change_time
				_allUserOrders[order.Taker][order.Tx_hash]["Trade_price"] = order.Trade_price
				//fmt.Println("scan buy order change time:", order.Taker, order.Order_change_time)
			}
			//fmt.Println(i, "processed")
		}
		//最后一页
		if ordersNum < perPage {
			break
		}
		offset += perPage
	} // <<<<<<<<<<<<<<<<<<<<< 结束 从 nftscan 读取全部成交定单 <<<<<<<<<<<<<<<<<<<<<
	return _allUserOrders
}

func calcPoints(allUserOrders map[string]map[string]map[string]interface{}, cfg *config.Config) map[string]int {

	UserPoints := make(map[string]int)
	for user, userOrder := range allUserOrders {
		for _, order := range userOrder {
			orderTxTime := fmtStrFromInterface(order["Order_change_time"])
			txTime, _ := strconv.Atoi(orderTxTime)
			if whiteListed(user, cfg.RankingConfig.WhiteList) {
				if txTime > cfg.RankingConfig.CbStart && txTime < cfg.RankingConfig.CbEnd {
					if order["Side"] == "sell" {
						UserPoints[user] = UserPoints[user] + cfg.RankingConfig.PointsPerSell*cfg.RankingConfig.CbTimes
					} else {
						UserPoints[user] = UserPoints[user] + cfg.RankingConfig.PointsPerBuy*cfg.RankingConfig.CbTimes
					}
				} else if txTime > cfg.RankingConfig.ObStart && txTime < cfg.RankingConfig.ObEnd {
					if order["Side"] == "sell" {
						UserPoints[user] = UserPoints[user] + cfg.RankingConfig.PointsPerSell
					} else {
						UserPoints[user] = UserPoints[user] + cfg.RankingConfig.PointsPerBuy
					}
				}
			} else {
				if txTime > cfg.RankingConfig.ObStart && txTime < cfg.RankingConfig.ObEnd {
					if order["Side"] == "sell" {
						UserPoints[user] = UserPoints[user] + cfg.RankingConfig.PointsPerSell
					} else {
						UserPoints[user] = UserPoints[user] + cfg.RankingConfig.PointsPerBuy
					}
				}
			}
		}
	}
	return UserPoints
}

func calcVolume(allUserOrders map[string]map[string]map[string]interface{}, cfg *config.Config) map[string]int {
	UserVolumes := make(map[string]int)
	for user, userOrder := range allUserOrders {
		for _, order := range userOrder {
			if order["Trade_price"] == "0" {
				continue
			}
			priceStr := fmtStrFromInterface(order["Trade_price"])
			priceInt64, err := strconv.ParseFloat(priceStr, 64)
			if nil == err {
				priceFloat := int(math.Trunc(priceInt64 / math.Pow10(12)))
				UserVolumes[user] += priceFloat
			} else {
				fmt.Println(err)
			}
		}
	}
	return UserVolumes
}

func writeUserTradePoints(userPoints map[string]int, keys []string, cfg *config.Config) {

	lenUserPoints := len(userPoints)
	userPointsArray := make([]map[string]interface{}, lenUserPoints)
	indexArray := 0
	for _, k := range keys {
		userPointsArray[indexArray] = make(map[string]interface{})
		userPointsArray[indexArray]["key"] = k
		userPointsArray[indexArray]["value"] = userPoints[k]
		userPointsArray[indexArray]["ranking"] = indexArray + 1
		indexArray++
	}
	str, err := json.Marshal(userPointsArray)
	if nil != err {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(cfg.RankingConfig.RankingFile, str, 0644)
	if nil != err {
		fmt.Println("failed to write ranking data file ")
	}
}

func writeUserVolumes(userVolumes map[string]int, userPoints map[string]int, keys []string, cfg *config.Config) {
	userVolumesArray := make([]map[string]interface{}, 100)
	indexArray := 0
	totalVolume := 0
	num := 0
	for _, k := range keys {
		totalVolume += userVolumes[k]
		num++
		if num >= 100 {
			break
		}
	}
	totalVolumeFloat := float64(totalVolume) / math.Pow10(6)
	totalPer := 0.0
	totalV := 0.0
	for _, k := range keys {
		userVolumesArray[indexArray] = make(map[string]interface{})
		userVolumesArray[indexArray]["key"] = k
		userVolumesArray[indexArray]["value"] = float64(userVolumes[k]) / math.Pow10(6)
		userVolumesArray[indexArray]["percent"] = (float64(userVolumes[k]) / math.Pow10(6)) * 100 / float64(totalVolumeFloat)
		userVolumesArray[indexArray]["ranking"] = indexArray + 1
		userVolumesArray[indexArray]["points"] = userPoints[k]
		totalPer += (float64(userVolumes[k]) / math.Pow10(6)) * 100 / float64(totalVolumeFloat)
		totalV += float64(userVolumes[k]) / math.Pow10(6)
		indexArray++
		if indexArray >= 100 {
			break
		}
	}
	fmt.Println("totalPer:", totalPer)
	fmt.Println("totalV:", totalV)
	fmt.Println("totalVolume:", totalVolumeFloat)
	str, err := json.Marshal((userVolumesArray))
	if nil == err {
		err := ioutil.WriteFile(cfg.RankingConfig.VolumeFile, str, 0644)
		if nil != err {
			fmt.Println("failed to write volumes file")
		}
	}
}
func fmtStrFromInterface(val interface{}) string {
	if val == nil {
		return ""
	}
	switch ret := val.(type) {
	case string:
		return ret
	case int:
		return fmt.Sprintf("%v", ret)
	}
	return ""
}

//是否在白名单内
func whiteListed(address string, whitelist []string) bool {
	for i := range whitelist {
		if whitelist[i] == address {
			return true
		}
	}
	return false
}

func isBlueChip(address string, bluechip []string) bool {
	for i := range bluechip {
		if bluechip[i] == address {
			return true
		}
	}
	return false
}

func clearListingMaps() {
	if len(allUserListingOrders) > 0 {
		for k := range allUserListingOrders {
			delete(allUserListingOrders, k)
		}
	}
}
func clearMaps() {
	if len(allScanMakerOrders) > 0 {
		for k := range allScanMakerOrders {
			delete(allScanMakerOrders, k)
		}
	}
	if len(allScanTakerOrders) > 0 {
		for k := range allScanTakerOrders {
			delete(allScanTakerOrders, k)
		}
	}

}

func readFromNftScan(offset int, readAll bool) ([]OrderInfo, error) {

	nftScanUrl := ""
	httpClient := &http.Client{}
	nftScanResult := NftScanResult{}

	if !readAll {
		nftScanUrl = fmt.Sprintf(nftScanApi, offset, perPage)
	} else {
		nftScanUrl = fmt.Sprintf(nftScanApiAll, offset, perPage)
	}
	req, err := http.NewRequest("GET", nftScanUrl, nil)
	if err != nil {
		fmt.Println(err)
		return nftScanResult.Data, err
	}
	req.Header.Add("X-API-KEY", xApiKey)
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nftScanResult.Data, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nftScanResult.Data, err
	}
	if err := json.Unmarshal(body, &nftScanResult); err != nil {
		return nftScanResult.Data, err
	}
	return nftScanResult.Data, nil
	//}
}
