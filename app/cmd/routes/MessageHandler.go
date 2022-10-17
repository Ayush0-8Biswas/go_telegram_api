package routes

import (
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go_telegram_api/app/cmd/config"
	"go_telegram_api/app/cmd/controllers"
)

func HandleMessage(mess *tgBotAPI.Message) error {
	if brg, ok := config.BridgingData[fmt.Sprint(mess.Chat.ID)]; ok {
		if len(mess.Photo) != 0 {
			return controllers.HandlePhoto(mess, brg.Destination)
		} else if mess.Video != nil {
			return controllers.HandleVideo(mess, brg.Destination)
		} else if mess.Document != nil {
			return controllers.HandleDocument(mess, brg.Destination)
		} else if mess.Audio != nil {
			return controllers.HandleAudio(mess, brg.Destination)
		} else if mess.Sticker != nil {
			return controllers.HandleSticker(mess, brg.Destination)
		} else if mess.Text != "" {
			return controllers.HandleText(mess, brg.Destination)
		}
	}
	return nil
}
