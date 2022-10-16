package routes

import (
	"github.com/gorilla/mux"
	"go_telegram_api/api/pkg/controllers"
)

var RegisterWApiRoutes = func(router *mux.Router) {
	router.HandleFunc("/sendMessage/", controllers.SendMessage).Methods("POST")
	router.HandleFunc("/sendPhoto/", controllers.SendPhoto).Methods("POST")
	router.HandleFunc("/sendAudio/", controllers.SendAudio).Methods("POST")
	router.HandleFunc("/sendDocument/", controllers.SendDocument).Methods("POST")
	router.HandleFunc("/sendVideo/", controllers.SendVideo).Methods("POST")
	router.HandleFunc("/sendContact/", controllers.SendContact).Methods("POST")
	router.HandleFunc("/sendSticker/", controllers.SendSticker).Methods("POST")
}
