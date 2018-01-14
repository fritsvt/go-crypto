package main

import(
	"github.com/labstack/echo"
	"github.com/PuerkitoBio/goquery"

	"strings"
	"encoding/json"
	"io/ioutil"
	"time"
	"net/http"
	"log"
	"fmt"
)

var SERVER_ADRESS = ":3000"
var INTERVAL time.Duration = 60 // interval new data should be fetched in seconds
var TARGET_URL = "https://coinmarketcap.com/all/views/all/"

func main() {
	e := echo.New()
	e.Static("/", "static")

	go loopInterval(INTERVAL) // start the loop

	type WelcomeMessage struct {
		Success bool `json:"success"`
		Message  string `json:"message"`
	}

	e.GET("/", func(c echo.Context) error {
		res := &WelcomeMessage{
			Success:true,
			Message:  "Welcome to the crypto-prices api, /coins.json for the data",
		}

		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(SERVER_ADRESS))
}

func scrapeCoins() {
	println("Fetching new coins...")

	type coin struct {
		Name string `json:"name"`
		Ticker string `json:"ticker"`
		Btc string `json:"btc"`
		Price string `json:"price"`
	}
	type responseStruct struct {
		Success bool `json:"success"`
		Timestamp int64 `json:"timestamp"`
		AmountOfCoins int `json:"amount_of_coins"`
		Coins []coin `json:"coins"`
	}

	var coins []coin
	var response[] responseStruct
	response = response;

	doc, err := goquery.NewDocument(TARGET_URL)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		id := s.AttrOr("id", "")
		if (id != "") {
			var price= s.Find(".price")

			var coinName = strings.TrimPrefix(id, "id-")
			var ticker = s.Find(".currency-symbol").Text()
			var btc = price.AttrOr("data-btc", "")
			var price_usd = price.AttrOr("data-usd", "")

			coins = append(coins, coin {
				Name:  coinName,
				Ticker: ticker,
				Btc: btc,
				Price: price_usd,
			})
		}
	})
	response = append(response, responseStruct{
		Success:true,
		Timestamp:time.Now().Unix(),
		AmountOfCoins:len(coins),
		Coins: coins,
	})

	//b, _ := json.MarshalIndent(response, "", " ")
	b, _ := json.MarshalIndent(response, "", " ")
	// writing json to file
	_ = ioutil.WriteFile("static/coins.json", b, 0644)

	println("Wrote " + fmt.Sprint(len(coins)) + " coins to the coins.json file")
}

func loopInterval(interval time.Duration) {
	scrapeCoins()
	for range time.Tick(time.Second * interval){
		scrapeCoins()
	}
}