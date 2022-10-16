package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"go_telegram_api/api/pkg/routes"
	"log"
	"net/http"
)

func APIStart(port string) {
	var r = mux.NewRouter()
	routes.RegisterWApiRoutes(r)
	http.Handle("/", r)
	fmt.Println("server started at http://localhost" + port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
