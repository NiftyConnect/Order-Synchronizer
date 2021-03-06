package adaptor

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/niftyConnect/order-synchronizer/config"
)

func TestExtractEvents(t *testing.T) {
	cfg := &config.Config{
		LogConfig: config.LogConfig{
			Level:                        "",
			Filename:                     "",
			MaxFileSizeInMB:              0,
			MaxBackupsOfLogFiles:         0,
			MaxAgeToRetainLogFilesInDays: 0,
			UseConsoleLogger:             false,
			UseFileLogger:                false,
			Compress:                     false,
		},
		ChainConfig: config.ChainConfig{
			ChainDetails: map[string]*config.ChainDetail{
				"BSC": {
					Providers: []string{
						"https://bsc-dataseed.binance.org",
						"https://bsc-dataseed1.defibit.io",
						"https://bsc-dataseed1.ninicoin.io",
					},
					StartHeight:      17706607,
					SleepWaitSecond:  3,
					ConfirmBlocks:    3,
					ExchangeContract: common.HexToAddress("0x643e5ae0B0464646C5Ba290eFD9aa31285450B83"),
				},
				"ETH": {
					Providers: []string{
						"wss://mainnet.infura.io/ws/v3/feed7cd0743b4d44abc4bcdac5086588",
						"wss://mainnet.infura.io/ws/v3/3b03dc6c49af444bacb681f07df43182",
					},
					StartHeight:      14754037,
					SleepWaitSecond:  10,
					ConfirmBlocks:    1,
					ExchangeContract: common.HexToAddress("0xaD7Ce050E176fDF23520330f0E7F342276b3c891"),
				},
			}},
		AlertConfig: config.AlertConfig{
			Identity:       "",
			TelegramBotId:  "",
			TelegramChatId: "",
		},
		DBConfig: config.DBConfig{
			Dialect: "sqlite3",
			DBPath:  "test.db",
		},
	}

	str, _ := json.Marshal(cfg)
	t.Log(string(str))
	//
	//adaptor, err := NewAdaptor(cfg, "ETH")
	//require.NoError(t, err)
	//
	//events, err := adaptor.ExtractEvents(1000, 14766550, 14784120)
	//require.NoError(t, err)
	//
	//for _, ev := range events {
	//	switch ev.(type) {
	//	case database.OrderApprovedPartOne:
	//		event := ev.(database.OrderApprovedPartOne)
	//		t.Log("OrderApprovedPartOne: " + event.Hash.String())
	//	case database.OrderApprovedPartTwo:
	//		event := ev.(database.OrderApprovedPartTwo)
	//		t.Log("OrderApprovedPartTwo: " + event.Hash.String())
	//	case database.OrderCancelled:
	//		event := ev.(database.OrderCancelled)
	//		t.Log("OrderCancelled: " + event.Hash.String())
	//	case database.OrdersMatched:
	//		event := ev.(database.OrdersMatched)
	//		t.Log("OrdersMatched: buyHash "+event.BuyHash.String()+" , sellHash ", event.SellHash.String())
	//	case database.NonceIncremented:
	//		event := ev.(database.NonceIncremented)
	//		t.Log("NonceIncremented: " + event.Maker.String())
	//	}
	//}
}
