package pricelist

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"topcoin/pricingservice/api"
	"topcoin/pricingservice/model"
)

// This is to handle the stocks that are not available in CoinMarket API
var invalidSymbols map[string]struct{} = make(map[string]struct{})

// Get Price List of all stocks
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

	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=" + stockString
	headers := map[string]string{"X-CMC_PRO_API_KEY": os.Getenv("API_KEY_PRICE")}
	body, err := api.GetData(url, headers)
	if err != nil {
		fmt.Println(err.Error())
	}

	parseErr := json.Unmarshal(body, &priceList)
	if parseErr != nil {
		log.Fatal(parseErr.Error())
	}

	errorString := priceList.Status.ErrorMessage
	if len(errorString) > 0 {

		// If error starts with "Index values for" it indicates that those stocks are not available.
		// All invalid symbols are removed from stock string and then again the API is called.
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
