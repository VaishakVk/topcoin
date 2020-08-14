package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"topcoin/clientservice/model"
)

func GetFromRatingService(limit int64) ([]model.PriceRanking, error) {
	var stockRanking model.APIResponse
	url := os.Getenv("RATING_SERVICE_URL") + "?limit=" + strconv.Itoa(int(limit))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	apiResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	parseErr := json.Unmarshal(apiResponse, &stockRanking)
	if parseErr != nil {
		return nil, errors.New(parseErr.Error())
	}
	fmt.Println(stockRanking)
	if !stockRanking.Status {
		return nil, errors.New(stockRanking.Error)
	}

	return stockRanking.Response, nil

}
