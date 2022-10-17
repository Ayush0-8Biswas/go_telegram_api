package connect

import (
	"encoding/json"
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go_telegram_api/app/cmd/config"
	"go_telegram_api/app/pkg/controllers"
	"go_telegram_api/app/pkg/models"
	"io/ioutil"
	"log"
)

func Connect(apiToken string, debug bool) {
	models.TgAPI, _ = tgBotAPI.NewBotAPI(apiToken)
	models.TgAPI.Debug = debug

	updateConfig := tgBotAPI.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := models.TgAPI.GetUpdatesChan(updateConfig)

	bridgeDataFile, err := ioutil.ReadFile("./app/cmd/config/bridge.json")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(bridgeDataFile, &config.BridgingData)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v", config.BridgingData)

	for update := range updates {
		go controllers.EventHandler(update)
	}
}
