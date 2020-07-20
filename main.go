package main

import (
	"goldnoti/api"
	"goldnoti/repository"
	"goldnoti/service"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func main() {
	service.SetupConfig("./config")
	service.SetupConfigParams()

	repository.SetupLINEConfig()
	repository.LINEBotInitialize()

	http.HandleFunc("/api/health", api.Health)
	http.HandleFunc("/api/today", api.GetTodayPrice)
	http.HandleFunc("/api/today/line", api.GetTodayPriceForLINEBot)

	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("Listening.Port")
	}
	http.ListenAndServe(":"+port, nil)
}
