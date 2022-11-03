package rankingserver

import (
	"encoding/json"
	"net/http"

	localcmm "github.com/niftyConnect/order-synchronizer/common"
)

func writeSuccessResponse(w http.ResponseWriter, responseData interface{}) {
	jsonBytes, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		localcmm.Logger.Errorf("write response error, err=%s", err.Error())
	}
}
