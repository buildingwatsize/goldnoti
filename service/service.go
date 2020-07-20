package service

import (
	"fmt"
	"goldnoti/model"
	"goldnoti/repository"
	"log"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
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

// HandleLINEEventMessage : Handle LINE Event Message for LINE Bot Client
func HandleLINEEventMessage(event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		textIncoming := message.Text
		log.Println("[INFO]: Text Incoming |", textIncoming)

		uid := event.Source.UserID
		log.Println("[INFO]: User ID |", uid)
		profile := repository.Bot.GetProfile(uid)
		log.Println("[INFO]: User Profile |", profile)

		todayPrice, err := repository.Harvester()
		if err != nil {
			log.Println("[ERROR]: Get Today Price Error |", err)
			return
		}

		flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(fmt.Sprintf(`
		{
			"type": "bubble",
			"hero": {
				"type": "image",
				"url": "https://lh3.googleusercontent.com/g7cB9uLxM-stNyzQfljbg_gX6MkLGBZa2RCJR8pB48xSxoKplB_Ag_6Gpg0WIB1-1iD4ZBZdnkALrz4_E6tL9sK_c_l1DnrlVJSWu9d9Ow5H-L7rJChl2wPlC4uVvdS-jypCcFBzdrB2V1Yprhs753G2w_yfyGBKDGbIXCuBWcP14hzMbLB81f-1tANMSEdX71DxLfwAgbSLPPqsn1TODCdfGttKuNSGKGhYzqCFHyltDRj15LhUQ6B4qALfoGIY7DyqqUrht295mOIt1mafEV7b2Vv15BnlyaZC4xeoL19bJJ2gR5Wy1_7Vp5iQ5WxTvmyetjYddeiazF-vBHSIaMrpHR7ilTSuF85-5RLcHRUEsthM6VD4ZUH8F7r3wfDiFgvU7I9CAY3oHT45uPMv8JX9iXTA3NezasR8nWGxxn8H0dG_qpm5Kv43U2LlY4IBn0takcV49xgwUxPGCa0027Ef3OBxsdHfisVfbQkQxA2HL7GShEG0Cmgx2wM_YdBZ-X3C3DoiPA8G4lKXgqoqORHmFO4D594TgfWaBEBKXyfbIbCNv1QqAG1ejjuS504nV8JMHgHDwFL8uouKc_1SAX1zJEkkcEjDUBsF4-F0QWsg3kChxfRMoJe_q89gdWpjrgiVjaFFFuXuOelfJyLhKI42_y5PUH-PVu-5WRUB4zbsJ_Mr5fDQIclvxBuZ=w1280-h704-no?authuser=0",
				"size": "full",
				"aspectRatio": "20:13",
				"aspectMode": "cover",
				"action": {
					"type": "uri",
					"label": "Action",
					"uri": "https://linecorp.com/"
				}
			},
			"body": {
				"type": "box",
				"layout": "vertical",
				"spacing": "md",
				"contents": [
					{
						"type": "text",
						"text": "ราคาทองวันนี้",
						"size": "xl",
						"gravity": "center",
						"weight": "bold",
						"wrap": true
					},
					{
						"type": "box",
						"layout": "vertical",
						"spacing": "sm",
						"margin": "lg",
						"contents": [
							{
								"type": "box",
								"layout": "baseline",
								"spacing": "sm",
								"contents": [
									{
										"type": "text",
										"text": "ทองคำแท่ง (ซื้อ)",
										"flex": 4,
										"size": "sm",
										"color": "#AAAAAA"
									},
									{
										"type": "text",
										"text": "%v",
										"flex": 3,
										"size": "sm",
										"align": "end",
										"color": "#666666",
										"wrap": true
									}
								]
							},
							{
								"type": "box",
								"layout": "baseline",
								"spacing": "sm",
								"contents": [
									{
										"type": "text",
										"text": "ทองคำแท่ง (ขาย)",
										"flex": 4,
										"size": "sm",
										"color": "#AAAAAA"
									},
									{
										"type": "text",
										"text": "%v",
										"flex": 3,
										"size": "sm",
										"align": "end",
										"color": "#666666",
										"wrap": true
									}
								]
							},
							{
								"type": "separator"
							},
							{
								"type": "box",
								"layout": "baseline",
								"spacing": "sm",
								"contents": [
									{
										"type": "text",
										"text": "ทองรูปพรรณ (ซื้อ)",
										"flex": 4,
										"size": "sm",
										"color": "#AAAAAA"
									},
									{
										"type": "text",
										"text": "%v",
										"flex": 3,
										"size": "sm",
										"align": "end",
										"color": "#666666",
										"wrap": true
									}
								]
							},
							{
								"type": "box",
								"layout": "baseline",
								"spacing": "sm",
								"contents": [
									{
										"type": "text",
										"text": "ทองรูปพรรณ (ขาย)",
										"flex": 4,
										"size": "sm",
										"color": "#AAAAAA"
									},
									{
										"type": "text",
										"text": "%v",
										"flex": 3,
										"size": "sm",
										"align": "end",
										"color": "#666666",
										"wrap": true
									}
								]
							},
							{
								"type": "separator"
							},
							{
								"type": "box",
								"layout": "baseline",
								"margin": "xxl",
								"contents": [
									{
										"type": "text",
										"text": "การเปลี่ยนแปลงวันนี้",
										"flex": 4,
										"size": "sm",
										"color": "#AAAAAA"
									},
									{
										"type": "text",
										"text": "%v (%v)",
										"flex": 3,
										"size": "sm",
										"align": "end",
										"weight": "bold",
										"color": "#666666",
										"wrap": true
									}
								]
							},
							{
								"type": "box",
								"layout": "vertical",
								"margin": "xxl",
								"contents": [
									{
										"type": "text",
										"text": "%v %v",
										"size": "sm",
										"align": "end"
									},
									{
										"type": "text",
										"text": "พัฒนาด้วย ❤️ จาก Goldnoti",
										"margin": "md",
										"size": "xs",
										"align": "end",
										"color": "#AAAAAA",
										"wrap": true
									}
								]
							}
						]
					}
				]
			}
		}
		`,
			fmt.Sprintf("%.2f", todayPrice.BarBuy),
			fmt.Sprintf("%.2f", todayPrice.BarSell),
			fmt.Sprintf("%.2f", todayPrice.OrnamentBuy),
			fmt.Sprintf("%.2f", todayPrice.OrnamentSell),
			fmt.Sprintf("%.2f", todayPrice.TodayChange),
			todayPrice.StatusChange,
			todayPrice.UpdatedDate,
			todayPrice.UpdatedTime,
		)))
		if err != nil {
			log.Println("[ERROR]: Flex Message Builder |", err)
			return
		}
		flexMessage := linebot.NewFlexMessage("ราคาทองคำวันนี้", flexContainer)

		replyToken := event.ReplyToken
		if _, err := repository.Bot.ReplyMessage(replyToken, flexMessage).Do(); err != nil {
			log.Println("[ERROR]: Reply Message |", err)
			return
		}
	}
}
