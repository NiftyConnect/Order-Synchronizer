package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	"github.com/niftyConnect/order-synchronizer/common"
	"github.com/niftyConnect/order-synchronizer/config"
)

type Server struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewServer(db *gorm.DB, cfg *config.Config) *Server {
	return &Server{
		db:  db,
		cfg: cfg,
	}
}

func (server *Server) Serve() {
	router := mux.NewRouter()

	router.HandleFunc(fmt.Sprintf("/"), server.handleQueryAPI).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/api/v1/query_order/{%s}/{%s}", queryParameterBlockchain, queryParameterOrderHash), server.handleQueryOrderByHash).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/api/v1/query_orders/{%s}/{%s}", queryParameterBlockchain, queryParameterNFTAddress), server.handleQueryOrderByNFTAddress).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/api/v1/query_orders/{%s}/{%s}/{%s}", queryParameterBlockchain, queryParameterNFTAddress, queryParameterTokenId), server.handleQueryOrderByNFTAddressAndTokenId).Methods("GET")

	listenAddr := defaultListenAddr
	if server.cfg.ServerConfig.ListenAddr != "" {
		listenAddr = server.cfg.ServerConfig.ListenAddr
	}
	srv := &http.Server{
		Handler:      cors.AllowAll().Handler(router),
		Addr:         listenAddr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	common.Logger.Infof("start server at %s", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("start server error, err=%s", err.Error()))
	}
}
