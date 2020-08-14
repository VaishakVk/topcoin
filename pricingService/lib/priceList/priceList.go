package pricelist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"topcoin/pricingservice/model"
)

var invalidSymbols map[string]struct{} = make(map[string]struct{})

func GetPriceList(stocks map[string]struct{}) model.PriceList {
	var priceList model.PriceList
	var stockString string

	// Loop through and generate string
	for key := range stocks {
		if _, ok := invalidSymbols[key]; !ok {
			stockString = stockString + key + ","
		}
	}
	stockString = strings.TrimRight(stockString, ",")

	fmt.Println("Stock String", stockString)
	request, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol="+stockString, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-CMC_PRO_API_KEY", os.Getenv("API_KEY_PRICE"))

	var client = http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	parseErr := json.Unmarshal(body, &priceList)
	if parseErr != nil {
		log.Fatal(parseErr.Error())
	}

	errorString := priceList.Status.ErrorMessage
	if len(errorString) > 0 {
		if strings.Index(errorString, "Invalid values for") != -1 {
			errorString = strings.ReplaceAll(errorString, `"`, ``)
			errorString = strings.ReplaceAll(errorString, `:`, `,`)
			invalidSymbolsArr := strings.Split(errorString, ",")
			for _, coin := range invalidSymbolsArr {
				invalidSymbols[strings.TrimSpace(coin)] = struct{}{}
			}
			priceList = GetPriceList(stocks)
		} else {
			fmt.Println(errorString)
		}
	}
	return priceList
}
