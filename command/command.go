package command

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/niftyConnect/order-synchronizer/common"
	"github.com/niftyConnect/order-synchronizer/config"
	"github.com/niftyConnect/order-synchronizer/database"
	"github.com/niftyConnect/order-synchronizer/ranking"
	"github.com/niftyConnect/order-synchronizer/rankingserver"
	"github.com/niftyConnect/order-synchronizer/server"
	"github.com/niftyConnect/order-synchronizer/synchronizer"
)

func readConfigAndInit() *config.Config {
	var cfg *config.Config
	configFilePath := viper.GetString(common.FlagConfigPath)
	if configFilePath == "" {
		panic("empty config path")
	}
	cfg = config.ParseConfigFromFile(configFilePath)

	common.InitLogger(&cfg.LogConfig)

	return cfg
}

func Start() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start order synchronizer",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := readConfigAndInit()

			var db *gorm.DB
			if cfg.DBConfig.DBPath != "" {
				var err error
				db, err = gorm.Open(cfg.DBConfig.Dialect, cfg.DBConfig.DBPath)
				if err != nil {
					panic(fmt.Sprintf("open db error, err=%s", err.Error()))
				}
				defer db.Close()
				database.InitTables(db)
			}

			for blockchain, chainDetails := range cfg.ChainConfig.ChainDetails {
				blockchain = config.FormatBlockchainName(blockchain)
				syncInst, err := synchronizer.NewSynchronizer(db, chainDetails, blockchain)
				if err != nil {
					panic(err)
				}
				syncInst.Start()
				common.Logger.Infof("Start order synchronizer for %s", blockchain)
			}

			serverInst := server.NewServer(db, cfg)
			serverInst.Serve()

			return nil
		},
	}
	return cmd
}
func MakeCloseBetaRanking() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ranking",
		Short: "make close beta ranking data",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := readConfigAndInit()
			go ranking.Start(cfg)
			go ranking.Listing(cfg)
			rankingServerInst := rankingserver.NewServer(cfg)
			rankingServerInst.Serve()

			return nil
		},
	}
	return cmd
}
