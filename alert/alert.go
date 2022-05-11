package alert

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/niftyConnect/order-synchronizer/common"
	"github.com/niftyConnect/order-synchronizer/config"
)

func SendTelegramMessage(cfg *config.Config, msg string) {
	if cfg.AlertConfig.TelegramBotId == "" || cfg.AlertConfig.TelegramChatId == "" || msg == "" {
		return
	}

	endPoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.AlertConfig.TelegramBotId)
	formData := url.Values{
		"chat_id":    {cfg.AlertConfig.TelegramChatId},
		"parse_mode": {"html"},
		"text":       {fmt.Sprintf("%s: %s", cfg.AlertConfig.Identity, msg)},
	}
	_, err := http.PostForm(endPoint, formData)
	if err != nil {
		common.Logger.Errorf(fmt.Sprintf("send telegram message error, msg=%s, err=%s", msg, err.Error()))
		return
	}
}
