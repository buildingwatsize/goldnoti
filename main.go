package main

import (
	"goldnoti/api"
	"goldnoti/service"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	service.SetupConfig("./config")
	service.SetupConfigParams()

	http.HandleFunc("/api/health", api.Health)
	http.HandleFunc("/api/today", api.GetTodayPrice)

	http.ListenAndServe(":"+viper.GetString("Listening.Port"), nil)
}
