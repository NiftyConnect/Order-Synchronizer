package config

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	LogConfig   LogConfig   `json:"log_config"`
	ChainConfig ChainConfig `json:"chain_config"`
	AlertConfig AlertConfig `json:"alert_config"`
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
	ExchangeContract common.Address `json:"exchange_contract"`
}

func ParseConfigFromJson(content string) *Config {
	var config Config
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}
	return &config
}
