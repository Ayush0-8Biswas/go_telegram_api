package cmd

import (
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go_telegram_api/app/pkg/models"
	"go_telegram_api/app/pkg/utils"
)

func Connect(apiToken string, debug bool) {
	models.TgAPI, _ = tgBotAPI.NewBotAPI(apiToken)
	models.TgAPI.Debug = debug

	updateConfig := tgBotAPI.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := models.TgAPI.GetUpdatesChan(updateConfig)

	for update := range updates {
		go func(update tgBotAPI.Update) {
			fmt.Println(utils.PrintUpdate(&update))
		}(update)
	}
}
