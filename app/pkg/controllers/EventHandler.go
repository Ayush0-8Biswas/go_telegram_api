package controllers

import (
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go_telegram_api/app/cmd/routes"
	"go_telegram_api/app/pkg/utils"
	"log"
)

func EventHandler(update tgBotAPI.Update) {
	if update.Message != nil {
		fmt.Println(update.Message.Chat.ID)
		err := routes.HandleMessage(update.Message)
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Println(utils.PrintUpdate(&update))
}
