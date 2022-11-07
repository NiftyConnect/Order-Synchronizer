package rankingserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	str := fmt.Sprintf("{\"data\":{\"CB_START\":%d,\"CB_END\":%d,\"OB_START\":%d, \"OB_END\":%d, \"PWAS\":%d, \"PWAB\":%d, \"CB_TIMES\":%d, \"MIN_LIST_TIME\":%d,\"wl\":%s,\"bc\":%s}}", server.cfg.RankingConfig.CbStart, server.cfg.RankingConfig.CbEnd, server.cfg.RankingConfig.ObStart, server.cfg.RankingConfig.ObEnd, server.cfg.RankingConfig.PointsPerSell, server.cfg.RankingConfig.PointsPerBuy, server.cfg.RankingConfig.CbTimes, server.cfg.RankingConfig.MinListedTime, wl, bc)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writeSuccessResponse(w, json.RawMessage(str))
}
