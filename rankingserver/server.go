package rankingserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/niftyConnect/order-synchronizer/common"
	"github.com/niftyConnect/order-synchronizer/config"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (server *Server) Serve() {
	fmt.Println("serving")
	router := mux.NewRouter()

	router.HandleFunc("/Ranking", server.handleRanking).Methods("GET")
	router.HandleFunc("/ranking", server.handleRanking).Methods("GET")
	router.HandleFunc("/params", server.handleGetParams).Methods("GET")
	router.HandleFunc("/twitterpoints", server.handlePostTwitterPoints).Methods("POST")
	router.HandleFunc("/gettwitterpoints", server.handleGetTwitterPoints).Methods("GET")
	router.HandleFunc("/getgeneralnft", server.handleGeneralNft).Methods("GET")
	router.HandleFunc("/volume", server.handleVolume).Methods("GET")

	router.HandleFunc("/getfile", server.handleGetFile).Methods("GET")

	router.HandleFunc("/claim", server.handleClaim).Methods("POST")
	router.HandleFunc("/getclaimednft", server.handleGetClaimedNft).Methods("GET")

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
