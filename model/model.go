package model

// Health : Struct Model for Health Check API
type Health struct {
	ProjectName      string `json:"project_name"`
	Status           string `json:"status"`
	Version          string `json:"version"`
	Env              string `json:"env"`
	RequestTimestamp string `json:"request_timestamp"`
}

// GoldPriceResponse : Data struct for response
type GoldPriceResponse struct {
	ResponseData      GoldPriceData `json:"response_data"`
	ResponseMessage   string        `json:"response_message"`
	ResponseTimestamp string        `json:"response_timestamp"`
}

// GoldPriceData : Data from Web Scraping
type GoldPriceData struct {
	BarBuy       float64 `json:"bar_buy"`
	BarSell      float64 `json:"bar_sell"`
	OrnamentBuy  float64 `json:"ornament_buy"`
	OrnamentSell float64 `json:"ornament_sell"`
	StatusChange string  `json:"status_change"`
	TodayChange  float64 `json:"today_change"`
	UpdatedDate  string  `json:"updated_date"`
	UpdatedTime  string  `json:"updated_time"`
}
