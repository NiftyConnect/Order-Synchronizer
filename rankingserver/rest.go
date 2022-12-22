package rankingserver

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func (server *Server) handleRanking(w http.ResponseWriter, r *http.Request) {
	var str, err = ioutil.ReadFile(server.cfg.RankingConfig.RankingFile)
	if nil != err {
		fmt.Println(err)
	}
	resp := fmt.Sprintf("{\"data\":%s}", str)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writeSuccessResponse(w, json.RawMessage(resp))
}

func (server *Server) handleGetParams(w http.ResponseWriter, r *http.Request) {
	wl, err := json.Marshal(server.cfg.RankingConfig.WhiteList)
	if nil != err {
		fmt.Println(err)
	}
	bc, err := json.Marshal(server.cfg.RankingConfig.BlueChip)
	if nil != err {
		fmt.Println(err)
	}
	tb, err := json.Marshal(server.cfg.RankingConfig.TB)
	if nil != err {
		fmt.Println(err)
	}
	str := fmt.Sprintf("{\"data\":{\"CB_START\":%d,\"CB_END\":%d,\"OB_START\":%d, \"OB_END\":%d, \"PWAS\":%d, \"PWAB\":%d, \"CB_TIMES\":%d, \"MIN_LIST_TIME\":%d,\"wl\":%s,\"bc\":%s,\"tb\":%s}}", server.cfg.RankingConfig.CbStart, server.cfg.RankingConfig.CbEnd, server.cfg.RankingConfig.ObStart, server.cfg.RankingConfig.ObEnd, server.cfg.RankingConfig.PointsPerSell, server.cfg.RankingConfig.PointsPerBuy, server.cfg.RankingConfig.CbTimes, server.cfg.RankingConfig.MinListedTime, wl, bc, tb)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writeSuccessResponse(w, json.RawMessage(str))
}

func (server *Server) handlePostTwitterPoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	result, _ := ioutil.ReadAll(r.Body)
	postData := make(map[string]string)
	json.Unmarshal(result, &postData)
	addresses := postData["addresses"]
	xApiKey := postData["xApiKey"]
	fmt.Println("xApiKey", xApiKey, "addresses:", addresses)
	errstr := ""
	if xApiKey != "xdkyywQi9ZoAHvg9DNeM" {
		errstr = fmt.Sprintf("Invalid api key:%s", xApiKey)
		writeSuccessResponse(w, errstr)
		return
	}

	addresses = strings.ToLower(addresses)
	filecontent, err := ioutil.ReadFile(twitterPointsFile)
	if err != nil {
		ioutil.WriteFile(twitterPointsFile, []byte(addresses), 0644)
		writeSuccessResponse(w, "OK")
	} else {
		if string(filecontent) == addresses {
			errstr = "List not changed"
			writeSuccessResponse(w, errstr)
			return
		} else {
			now := time.Now()
			newFile := fmt.Sprintf("%d%02d%02d-%02d%02d%02d%s",
				now.Year(), now.Month(), now.Day(),
				now.Hour(), now.Minute(), now.Second(),
				string(twitterPointsFile),
			)
			ioutil.WriteFile(newFile, filecontent, 0644)
			ioutil.WriteFile(twitterPointsFile, []byte(addresses), 0644)
			writeSuccessResponse(w, "OK")
		}
	}
}

func (server *Server) handleGetTwitterPoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()
	action := r.FormValue("action")
	if action == "history" {
		fileSystem := os.DirFS(".")
		retstr := ""
		i := 0
		fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
			if i >= 20 {
				return nil
			}
			if err != nil {
				fmt.Println(err)
			}
			if strings.Contains(path, twitterPointsFile) {
				if strings.Index(path, twitterPointsFile) != 0 {
					retstr = retstr + strings.Replace(path, twitterPointsFile, "", -1) + "\n"
					i++
				}
			}
			return nil
		})
		writeSuccessResponse(w, string(retstr))

	} else {
		filecontent, _ := ioutil.ReadFile(twitterPointsFile)
		writeSuccessResponse(w, string(filecontent))
	}
}

func (server *Server) handleGetFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()
	fn := r.FormValue("file")
	retstr, err := ioutil.ReadFile(fn + twitterPointsFile)
	if err != nil {
		fmt.Println(err)
	} else {
		writeSuccessResponse(w, string(retstr))
	}

}

func (server *Server) handleGeneralNft(w http.ResponseWriter, r *http.Request) {
	var str, err = ioutil.ReadFile(server.cfg.RankingConfig.NftFile)
	if nil != err {
		fmt.Println(err)
	}
	resp := fmt.Sprintf("{\"data\":%s}", str)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writeSuccessResponse(w, json.RawMessage(resp))

}

func (server *Server) handleVolume(w http.ResponseWriter, r *http.Request) {
	var str, err = ioutil.ReadFile(server.cfg.RankingConfig.VolumeFile)
	if nil != err {
		fmt.Println(err)
	}
	resp := fmt.Sprintf("{\"data\":%s}", str)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writeSuccessResponse(w, json.RawMessage(resp))

}

func (server *Server) handleClaim(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	result, _ := ioutil.ReadAll(r.Body)
	postData := make(map[string]string)
	claimFileData := make(map[string]map[string]int)
	json.Unmarshal(result, &postData)
	address := postData["address"]
	nftType := postData["nftType"]
	fmt.Println("address:", address, "nftType:", nftType)
	if len(address) < 20 {
		writeSuccessResponse(w, "missing address")
		return
	}
	if len(nftType) < 1 {
		writeSuccessResponse(w, "missing nftType")
		return
	}
	if (nftType != "general") && (nftType != "premium") {
		writeSuccessResponse(w, "invalid nftType")
	} else {
		var old_content, err = ioutil.ReadFile(server.cfg.RankingConfig.ClaimFile)
		if err != nil {
			writeSuccessResponse(w, err)
		} else {
			if len(old_content) == 0 {
				claimFileData[address] = make(map[string]int)
				claimFileData[address][nftType] = 1
			} else {
				err := json.Unmarshal(old_content, &claimFileData)
				if nil != err {
					fmt.Println(err)
				} else {
					if claimFileData[address] == nil {
						claimFileData[address] = make(map[string]int)
					} else {
						if claimFileData[address][nftType] == 1 {
							writeSuccessResponse(w, "already claimed")
							return
						}
					}
					claimFileData[address][nftType] = 1
				}
			}
			new_content, err := json.Marshal(claimFileData)
			if err != nil {
				writeSuccessResponse(w, err)
			}
			ioutil.WriteFile(server.cfg.RankingConfig.ClaimFile, []byte(new_content), 0644)
			writeSuccessResponse(w, "ok")
		}
	}
}

func (server *Server) handleGetClaimedNft(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()
	address := r.FormValue("address")
	var content, err = ioutil.ReadFile(server.cfg.RankingConfig.ClaimFile)
	var fileData = make(map[string]map[string]int)
	if err != nil {
		writeSuccessResponse(w, err)
		return
	} else {
		if len(content) > 0 && len(address) > 1 {
			err := json.Unmarshal(content, &fileData)
			if err != nil {
				writeSuccessResponse(w, err)
				return
			}
			msg, err := json.Marshal(fileData[address])
			if err != nil {
				fmt.Println(err)
			}
			writeSuccessResponse(w, json.RawMessage(msg))
		} else {
			writeSuccessResponse(w, json.RawMessage(content))
		}
	}
}
