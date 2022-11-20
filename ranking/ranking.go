package ranking

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

func Start(cfg *config.Config) {

	//隔一段时间重新跑
	for {
		offset := 0
		clearMaps()
		t1 := time.Now().Unix()
		//fmt.Println("Rebuild Ranking data...")
		//读取全部成交定单
		for {
			scanOrders, err := readFromNftScan(offset)

			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("retry in ", retryInterval, "seconds")
				time.Sleep(time.Duration(retryInterval) * time.Second)
				continue
			}
			ordersNum := len(scanOrders)
			fmt.Println("len:", ordersNum)

			for _, order := range scanOrders {

				if nil == allUserOrders[order.Maker] {
					//fmt.Println("add maker map")
					allUserOrders[order.Maker] = make(map[string]map[string]interface{})
				}
				//fmt.Println(allUserOrders[order.Maker][order.Tx_hash])
				if len(allUserOrders[order.Maker][order.Tx_hash]) == 0 {
					//fmt.Println("add maker hash => order")
					allUserOrders[order.Maker][order.Tx_hash] = make(map[string]interface{})
					allUserOrders[order.Maker][order.Tx_hash]["Side"] = "sell"
					allUserOrders[order.Maker][order.Tx_hash]["Maker"] = order.Maker
					allUserOrders[order.Maker][order.Tx_hash]["Taker"] = order.Taker
					allUserOrders[order.Maker][order.Tx_hash]["Tx_time"] = order.Tx_time
					allUserOrders[order.Maker][order.Tx_hash]["Tx_hash"] = order.Tx_hash
					allUserOrders[order.Maker][order.Tx_hash]["Token_id"] = order.Token_id
					allUserOrders[order.Maker][order.Tx_hash]["Nft_address"] = order.Nft_address
					allUserOrders[order.Maker][order.Tx_hash]["Listing_time"] = order.Listing_time
					allUserOrders[order.Maker][order.Tx_hash]["Order_change_time"] = order.Order_change_time
					//fmt.Println("scan sell order change time:", order.Order_change_time)

				}
				if nil == allUserOrders[order.Taker] {
					//fmt.Println("add taker map")
					allUserOrders[order.Taker] = make(map[string]map[string]interface{})
				}

				if len(allUserOrders[order.Taker][order.Tx_hash]) == 0 {
					//fmt.Println("add taker hask => order")
					allUserOrders[order.Taker][order.Tx_hash] = make(map[string]interface{})
					allUserOrders[order.Taker][order.Tx_hash]["Side"] = "buy"
					allUserOrders[order.Taker][order.Tx_hash]["Maker"] = order.Maker
					allUserOrders[order.Taker][order.Tx_hash]["Taker"] = order.Taker
					allUserOrders[order.Taker][order.Tx_hash]["Tx_time"] = order.Tx_time
					allUserOrders[order.Taker][order.Tx_hash]["Tx_hash"] = order.Tx_hash
					allUserOrders[order.Taker][order.Tx_hash]["Token_id"] = order.Token_id
					allUserOrders[order.Taker][order.Tx_hash]["Nft_address"] = order.Nft_address
					allUserOrders[order.Taker][order.Tx_hash]["Listing_time"] = order.Listing_time
					allUserOrders[order.Taker][order.Tx_hash]["Order_change_time"] = order.Order_change_time
					//fmt.Println("scan buy order change time:", order.Taker, order.Order_change_time)
				}
				//fmt.Println(i, "processed")
			}
			//最后一页
			if ordersNum < perPage {
				break
			}
			offset += perPage
		}

		//计算排行榜
		//清空
		for i := range userPoints {
			userPoints[i] = 0
		}

		for i := range userPointsSort {
			userPointsSort[i] = 0
		}
		for user, userOrder := range allUserOrders {
			for _, order := range userOrder {
				orderTxTime := fmtStrFromInterface(order["Order_change_time"])
				txTime, _ := strconv.Atoi(orderTxTime)
				//fmt.Println(">>> ========================== New Order ==========================")
				//f//mt.Print("[", user, "] ")
				//白名单用户
				if whiteListed(user, cfg.RankingConfig.WhiteList) {
					//fmt.Println("whitelisted ")
					//fmt.Println("txTime:", txTime)
					//fmt.Println("side:", order["Side"])
					//fmt.Println("changetime:", order["Order_change_time"], txTime)
					//close beta 期间完成
					if txTime > cfg.RankingConfig.CbStart && txTime < cfg.RankingConfig.CbEnd {
						//fmt.Print("During CB, ")
						if order["Side"] == "sell" {
							//fmt.Println(" made a deal as [SELLER] won [", cfg.RankingConfig.PointsPerSell*cfg.RankingConfig.CbTimes, "] points")
							userPoints[user] = userPoints[user] + cfg.RankingConfig.PointsPerSell*cfg.RankingConfig.CbTimes
						} else {
							//fmt.Println(" made a deal as [BUYER] won [", cfg.RankingConfig.PointsPerBuy*cfg.RankingConfig.CbTimes, "] points")
							userPoints[user] = userPoints[user] + cfg.RankingConfig.PointsPerBuy*cfg.RankingConfig.CbTimes
						}
					} else if txTime > cfg.RankingConfig.ObStart && txTime < cfg.RankingConfig.ObEnd {
						//fmt.Print("During OB, ")
						if order["Side"] == "sell" {
							//fmt.Println(" made a deal as [SELLER] won [", cfg.RankingConfig.PointsPerSell, "] points")
							userPoints[user] = userPoints[user] + cfg.RankingConfig.PointsPerSell
						} else {
							//fmt.Println(" made a deal as [BUYER] won [", cfg.RankingConfig.PointsPerBuy, "] points")
							userPoints[user] = userPoints[user] + cfg.RankingConfig.PointsPerBuy
						}
					}
				} else {
					//fmt.Println("not whitelisted")
					if txTime > cfg.RankingConfig.ObStart && txTime < cfg.RankingConfig.ObEnd {
						//fmt.Print("During OB, ")
						if order["Side"] == "sell" {
							//fmt.Println(" made a deal as [SELLER] won [", cfg.RankingConfig.PointsPerSell, "] points")
							userPoints[user] = userPoints[user] + cfg.RankingConfig.PointsPerSell
						} else {
							//fmt.Println(" made a deal as [BUYER] won [", cfg.RankingConfig.PointsPerBuy, "] points")
							userPoints[user] = userPoints[user] + cfg.RankingConfig.PointsPerBuy
						}
					}
				}
				//fmt.Println("<<< ========================== Order End ==========================")
				//fmt.Println("")
			}
		}

		//twitter
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

		//排序
		keys := make([]string, 0, len(userPoints))
		for key := range userPoints {
			keys = append(keys, key)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			return userPoints[keys[i]] > userPoints[keys[j]]
		})
		lenUserPoints := len(userPoints)
		userPointsArray := make([]map[string]interface{}, lenUserPoints)
		indexArray := 0
		for _, k := range keys {
			userPointsArray[indexArray] = make(map[string]interface{})
			userPointsArray[indexArray]["key"] = k
			userPointsArray[indexArray]["value"] = userPoints[k]
			indexArray++
		}

		/* 		for k, v := range userPointsSort {
			userPointsArray[indexArray] = make(map[string]interface{})
			userPointsArray[indexArray]["key"] = k
			userPointsArray[indexArray]["value"] = v
			indexArray++
			fmt.Println("value:", v)
		} */

		str, err := json.Marshal(userPointsArray)
		if nil != err {
			fmt.Println(err)
		}
		//fmt.Println(string(str))
		err = ioutil.WriteFile(cfg.RankingConfig.RankingFile, str, 0644)
		if nil != err {
			fmt.Println("failed to write ranking data file ")
		}
		t2 := time.Now().Unix()
		fmt.Println("It took", (t2 - t1), "seconds to make ranking")
		time.Sleep(time.Duration(cfg.RankingConfig.RankingInterval) * time.Second)
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

//是否是蓝筹
/* func isBlueChip(address string, bluechip []string) bool {
	return false
} */

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

func readFromNftScan(offset int) ([]OrderInfo, error) {

	nftScanUrl := ""
	httpClient := &http.Client{}
	nftScanResult := NftScanResult{}
	fmt.Println("read page [", (1 + offset/perPage), "] from nftscan, perPage:", perPage)
	//for {
	nftScanUrl = fmt.Sprintf(nftScanApi, offset, perPage)
	fmt.Println(nftScanUrl)
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
	//io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nftScanResult.Data, err
	}
	//bodyStr := string(body)
	if err := json.Unmarshal(body, &nftScanResult); err != nil {
		return nftScanResult.Data, err
		//fmt.Println(bodyStr)
	}
	//fmt.Printf("%s", fmt.Sprintf("nftscan return code:%d  msg:%s, len:%d", nftScanResult.Code, nftScanResult.Msg, len(nftScanResult.Data)))
	//fmt.Println("")
	return nftScanResult.Data, nil
	//}
}
