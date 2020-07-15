package api

import (
	"encoding/json"
	"goldnoti/model"
	"goldnoti/service"
	"log"
	"net/http"
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
	methodNotAllowed := "Method Not Allowed"
	log.Println(methodNotAllowed)
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(model.GoldPriceResponse{
		ResponseData:      model.GoldPriceData{},
		ResponseMessage:   methodNotAllowed,
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
