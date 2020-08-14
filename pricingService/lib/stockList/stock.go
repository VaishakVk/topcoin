package stocklist

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"topcoin/pricingservice/api"
	"topcoin/pricingservice/model"
)

// GetAvialbleStocks - get all available stocks
func GetAvailableStocks() model.CoinList {
	var coinList model.CoinList

	apiKey := os.Getenv("API_KEY_LIST")
	url := "https://min-api.cryptocompare.com/data/blockchain/list?api_key=" + apiKey
	resp, err := api.GetData(url, nil)
	if err != nil {
		fmt.Println(err)
	}
	parseErr := json.Unmarshal(resp, &coinList)
	if parseErr != nil {
		log.Fatal(parseErr.Error())
	}
	return coinList
}
