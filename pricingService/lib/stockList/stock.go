package stocklist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"topcoin/pricingservice/model"
)

func GetAvailableStocks() model.CoinList {
	var coinList model.CoinList

	apiKey := os.Getenv("API_KEY_LIST")
	url := "https://min-api.cryptocompare.com/data/blockchain/list?api_key=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	apiResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	parseErr := json.Unmarshal(apiResponse, &coinList)
	if parseErr != nil {
		log.Fatal(parseErr.Error())
	}
	return coinList
}
