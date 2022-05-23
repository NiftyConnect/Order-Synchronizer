package server

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"

	"github.com/niftyConnect/order-synchronizer/config"
	"github.com/niftyConnect/order-synchronizer/database"
)

func (server *Server) handleQueryAPI(w http.ResponseWriter, r *http.Request) {
	apis := []string{
		fmt.Sprintf("/api/v1/query_order/{%s}/{%s}", queryParameterBlockchain, queryParameterOrderHash),
		fmt.Sprintf("/api/v1/query_orders/{%s}/{%s}", queryParameterBlockchain, queryParameterNFTAddress),
		fmt.Sprintf("/api/v1/query_orders/{%s}/{%s}/{%s}", queryParameterBlockchain, queryParameterNFTAddress, queryParameterTokenId),
	}

	writeSuccessResponse(w, apis)
}

func (server *Server) handleQueryOrderByHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockchain := config.FormatBlockchainName(vars[queryParameterBlockchain])
	orderHash := vars[queryParameterOrderHash]

	var order database.NiftyConnectOrder
	if err := server.db.Model(database.NiftyConnectOrder{}).
		Where("blockchain = ? and order_hash = ?", blockchain, orderHash).First(&order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeSuccessResponse(w, order)
}

func (server *Server) handleQueryOrderByNFTAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockchain := config.FormatBlockchainName(vars[queryParameterBlockchain])
	nftAddress := common.HexToAddress(vars[queryParameterNFTAddress]).String()

	query := r.URL.Query()
	offsetStr := query.Get("offset")
	limitStr := query.Get("limit")

	var err error
	offset := 0
	if len(offsetStr) != 0 {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid offset err: %s", err.Error()), http.StatusBadRequest)
			return
		}
	}
	limit := defaultLimit
	if len(limitStr) != 0 {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid limit size err: %s", err.Error()), http.StatusBadRequest)
			return
		}
		if limit == 0 || limit > defaultLimit {
			limit = defaultLimit
		}
	}

	var orders []database.NiftyConnectOrder
	if err := server.db.Model(database.NiftyConnectOrder{}).
		Where("blockchain = ? and nft_address = ?", blockchain, nftAddress).Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeSuccessResponse(w, orders)
}

func (server *Server) handleQueryOrderByNFTAddressAndTokenId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockchain := config.FormatBlockchainName(vars[queryParameterBlockchain])
	nftAddress := common.HexToAddress(vars[queryParameterNFTAddress]).String()
	tokenIdStr := vars[queryParameterTokenId]

	tokenId := big.NewInt(0)
	if strings.HasPrefix(tokenIdStr, "0x") {
		_, ok := tokenId.SetString(tokenIdStr, 16)
		if !ok {
			http.Error(w, fmt.Sprintf("invalid tokenId %s", tokenIdStr), http.StatusBadRequest)
			return
		}
	} else {
		_, ok := tokenId.SetString(tokenIdStr, 10)
		if !ok {
			http.Error(w, fmt.Sprintf("invalid tokenId %s", tokenIdStr), http.StatusBadRequest)
			return
		}
	}

	query := r.URL.Query()
	offsetStr := query.Get("offset")
	limitStr := query.Get("limit")

	var err error
	offset := 0
	if len(offsetStr) != 0 {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid offset err: %s", err.Error()), http.StatusBadRequest)
			return
		}
	}
	limit := defaultLimit
	if len(limitStr) != 0 {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid limit size err: %s", err.Error()), http.StatusBadRequest)
			return
		}
		if limit == 0 || limit > defaultLimit {
			limit = defaultLimit
		}
	}

	var orders []database.NiftyConnectOrder
	if err := server.db.Model(database.NiftyConnectOrder{}).
		Where("blockchain = ? and nft_address = ? and token_id = ?", blockchain, nftAddress, tokenId.String()).Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeSuccessResponse(w, orders)
}
