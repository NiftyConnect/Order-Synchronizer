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
