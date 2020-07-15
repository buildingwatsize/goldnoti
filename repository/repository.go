package repository

import (
	"errors"
	"goldnoti/model"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const (
	// GoldtradersLink : link for ทองคำราคา.com
	GoldtradersLink = "https://xn--42cah7d0cxcvbbb9x.com/"
	// GoldtradersTableElement : parent element for pricing zone
	GoldtradersTableElement = "#rightCol > div.divgta.goldshopf > table > tbody"
)

// ToFloat64 : String Converter for Float64
func ToFloat64(stringVal string) float64 {
	stringValNoComma := strings.ReplaceAll(stringVal, ",", "")
	floatVal, err := strconv.ParseFloat(stringValNoComma, 64)
	if err != nil {
		return 0.00
	}
	return floatVal
}

// Harvester : Core Func for harvesting data, I use Colly for Web Scraping
func Harvester() (model.GoldPriceData, error) {
	var (
		barSell      string
		barBuy       string
		ornamentSell string
		ornamentBuy  string
		statusChange string
		todayChange  string
		updatedDate  string
		updatedTime  string
	)

	c := colly.NewCollector()

	c.OnHTML(GoldtradersTableElement, func(e *colly.HTMLElement) {
		barSell = e.ChildText("tr:nth-child(2) > td:nth-child(2)")
		barBuy = e.ChildText("tr:nth-child(2) > td:nth-child(3)")
		ornamentSell = e.ChildText("tr:nth-child(3) > td:nth-child(2)")
		ornamentBuy = e.ChildText("tr:nth-child(3) > td:nth-child(3)")
		statusChange = e.ChildAttr("tr:nth-child(4) > td.span > .imgp", "alt")
		todayChange = strings.Split(e.ChildText("tr:nth-child(4) > td:nth-child(1)"), " ")[1]
		updatedDate = e.ChildText("tr:nth-child(5) > td.span.bg-span.txtd.al-r")
		updatedTime = e.ChildText("tr:nth-child(5) > td.em.bg-span.txtd.al-r")
	})

	c.Visit(GoldtradersLink)

	log.Println("ทองแท่ง ขายออก:", barSell)
	log.Println("ทองแท่ง รับซื้อ:", barBuy)
	log.Println("ทองรูปพรรณ ขายออก:", ornamentSell)
	log.Println("ทองรูปพรรณ รับซื้อ:", ornamentBuy)
	log.Println("สถานะ:", statusChange)
	log.Println("เปลี่ยนแปลงวันนี้:", todayChange)
	log.Println("วันที่อัพเดท:", updatedDate)
	log.Println("เวลาที่อัพเดท:", updatedTime)

	if barSell == "" || barBuy == "" || ornamentSell == "" || ornamentBuy == "" || statusChange == "" || todayChange == "" || updatedDate == "" || updatedTime == "" {
		return model.GoldPriceData{}, errors.New("Not found data")
	}

	return model.GoldPriceData{
		BarBuy:       ToFloat64(barSell),
		BarSell:      ToFloat64(barBuy),
		OrnamentBuy:  ToFloat64(ornamentSell),
		OrnamentSell: ToFloat64(ornamentBuy),
		StatusChange: statusChange,
		TodayChange:  ToFloat64(todayChange),
		UpdatedDate:  updatedDate,
		UpdatedTime:  updatedTime,
	}, nil
}
