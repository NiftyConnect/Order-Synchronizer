package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	LogConfig     LogConfig     `json:"log_config"`
	ChainConfig   ChainConfig   `json:"chain_config"`
	AlertConfig   AlertConfig   `json:"alert_config"`
	DBConfig      DBConfig      `json:"db_config"`
	ServerConfig  ServerConfig  `json:"server_config"`
	RankingConfig RankingConfig `json:"ranking_config"`
}

type AlertConfig struct {
	Identity       string `json:"identity"`
	TelegramBotId  string `json:"telegram_bot_id"`
	TelegramChatId string `json:"telegram_chat_id"`
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

type ChainConfig struct {
	ChainDetails map[string]*ChainDetail `json:"chain_details"`
}

type ChainDetail struct {
	Providers        []string       `json:"providers"`
	StartHeight      int64          `json:"start_height"`
	SleepWaitSecond  int64          `json:"sleep_wait_second"`
	ConfirmBlocks    int64          `json:"confirm_blocks"`
	HeightStep       int64          `json:"height_step"`
	ExchangeContract common.Address `json:"exchange_contract"`
}

type DBConfig struct {
	Dialect string `json:"dialect"`
	DBPath  string `json:"db_path"`
}

type ServerConfig struct {
	ListenAddr string `json:"listen_addr"`
}

type RankingConfig struct {
	RankingFile     string   `json:"ranking_file"`
	NftFile         string   `json:"nft_file"`
	VolumeFile      string   `json:"volume_file"`
	ClaimFile       string   `json:"claim_file"`
	WhiteList       []string `json:"whitelist"`
	BlueChip        []string `json:"bluechip"`
	TB              []string `json:"tb"`
	CbStart         int      `json:"cbstart"`
	CbEnd           int      `json:"cbend"`
	ObStart         int      `json:"obstart"`
	ObEnd           int      `json:"obend"`
	PointsPerSell   int      `json:"points_per_sell"`
	PointsPerBuy    int      `json:"points_per_buy"`
	CbTimes         int      `json:"cb_times"`
	RankingInterval int      `json:"ranking_interval"`
	MinListedTime   int      `json:"min_listed_time"`
}

func ParseConfigFromFile(filePath string) *Config {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	return &config
}
