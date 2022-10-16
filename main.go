package main

import (
	"github.com/joho/godotenv"
	cmd2 "go_telegram_api/api/cmd"
	"go_telegram_api/app/cmd"
)

func main() {
	var env, _ = godotenv.Read(".env")
	var apiToken = env["botAPI"]
	go cmd.Connect(apiToken, false)
	cmd2.APIStart(":9000")
}
