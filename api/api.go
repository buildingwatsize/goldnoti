package api

import (
	"encoding/json"
	"goldnoti/model"
	"goldnoti/repository"
	"goldnoti/service"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json"
)

// Health : Health Check API
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeKey, contentTypeValue)
	serviceData := service.HealthCheck()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serviceData)
}

// isPOST : checking function for request, it is POST or not!?
func isPOST(r *http.Request) bool {
	log.Println("Incoming Method:", r.Method)
	return r.Method == "POST"
}

// ResponseMethodNotAllowed : Response to requester in case 405: Method Not Allowed
func ResponseMethodNotAllowed(w http.ResponseWriter) {
	responseMessage := "Method Not Allowed"
	log.Println(responseMessage)
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(model.GoldPriceResponse{
		ResponseData:      model.GoldPriceData{},
		ResponseMessage:   responseMessage,
		ResponseTimestamp: service.GetTimestamp(),
	})
}

// ResponseBadRequest : Response to requester in case 400: Bad Request
func ResponseBadRequest(w http.ResponseWriter) {
	responseMessage := "Bad Request"
	log.Println(responseMessage)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(model.GoldPriceResponse{
		ResponseData:      model.GoldPriceData{},
		ResponseMessage:   responseMessage,
		ResponseTimestamp: service.GetTimestamp(),
	})
}

// ResponseInternalServerError : Response to requester in case 500: Internal Server Error
func ResponseInternalServerError(w http.ResponseWriter) {
	responseMessage := "Internal Server Error"
	log.Println(responseMessage)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(model.GoldPriceResponse{
		ResponseData:      model.GoldPriceData{},
		ResponseMessage:   responseMessage,
		ResponseTimestamp: service.GetTimestamp(),
	})
}

// GetTodayPrice : Get today gold price
func GetTodayPrice(w http.ResponseWriter, r *http.Request) {
	log.Println("------- API: GetTodayPrice -------")
	w.Header().Set(contentTypeKey, contentTypeValue)
	if !isPOST(r) {
		ResponseMethodNotAllowed(w)
		return
	}
	todayPrice := service.GetTodayPrice()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todayPrice)
}

// GetTodayPriceForLINEBot : Get today gold price with webhook
func GetTodayPriceForLINEBot(w http.ResponseWriter, r *http.Request) {
	log.Println("------- API: GetTodayPriceForLINEBot -------")
	defer r.Body.Close()
	w.Header().Set(contentTypeKey, contentTypeValue)
	if !isPOST(r) {
		ResponseMethodNotAllowed(w)
		return
	}

	events, err := repository.Bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ResponseBadRequest(w)
		} else {
			ResponseInternalServerError(w)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			go service.HandleLINEEventMessage(event)
		}
	}
}
