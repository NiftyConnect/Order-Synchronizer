package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/niftyConnect/order-synchronizer/database"
)

func (server *Server) handleQueryOrderByHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockchain := vars[queryParameterBlockchain]
	orderHash := vars[queryParameterOrderHash]

	var order database.NiftyConnectOrder
	if err := server.db.Model(database.NiftyConnectOrder{}).
		Where("blockchain = ? and order_hash = ?", blockchain, orderHash).First(&order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeSuccessResponse(w, order)
}
