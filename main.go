package main

import (
	"github.com/joho/godotenv"
	cmd2 "go_telegram_api/api/cmd"
	"go_telegram_api/app/cmd/config"
	"go_telegram_api/app/cmd/connect"
)

func main() {
	var env, _ = godotenv.Read(".env")
	var apiToken = env["botAPI"]
	//fmt.Printf("%#v", apiToken)
	go connect.Connect(apiToken, false)
	cmd2.APIStart(config.TelegramPort)
}
