package service

import (
	"fmt"
	"goldnoti/model"
	"goldnoti/repository"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	timestampFormat = "2006-01-02T15:04:05Z"
)

var (
	// TimeZone : Bangkok Thailand
	TimeZone string
)

// SetupConfig : reading configuration file via Viper
func SetupConfig(configPath string) {
	viper.SetConfigType("json")
	viper.SetConfigName("config." + os.Getenv("ENV"))
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	viper.Debug()
}

// SetupConfigParams : setting up configuration
func SetupConfigParams() {
	TimeZone = viper.GetString("TimeZone")
}

// HealthCheck : for get current health status
func HealthCheck() model.Health {
	return model.Health{
		ProjectName:      viper.GetString("ProjectName"),
		Status:           "I'm OK.",
		Version:          viper.GetString("Version"),
		Env:              viper.GetString("ENV"),
		RequestTimestamp: GetTimestamp(),
	}
}

// GetTimestamp : for get current timestamp
func GetTimestamp() string {
	date := time.Now()
	location, err := time.LoadLocation(TimeZone)
	if err != nil {
		log.Println("Get Timestamp Error ", err)
		return date.Format(timestampFormat)
	}
	return date.In(location).Format(timestampFormat)
}

// GetTodayPrice : Get Today Gold Price Service
func GetTodayPrice() model.GoldPriceResponse {
	var dataReturn model.GoldPriceResponse

	dataFromHarvested, err := repository.Harvester()
	if err != nil {
		dataReturn.ResponseMessage = err.Error()
	}

	dataReturn.ResponseData = dataFromHarvested
	dataReturn.ResponseTimestamp = GetTimestamp()

	return dataReturn
}
