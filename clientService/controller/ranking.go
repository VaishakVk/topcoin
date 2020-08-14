package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"topcoin/clientservice/api"
	"topcoin/clientservice/response"
)

// API Handler to get top coins
func GetTopCoins(res http.ResponseWriter, req *http.Request) {
	keys, exist := req.URL.Query()["limit"]
	if !exist {
		response.SendResponse(res, http.StatusBadRequest, false, nil, "Limit is invalid")
		return
	}

	limit, err := strconv.ParseInt(keys[0], 10, 32)
	if err != nil {
		response.SendResponse(res, http.StatusBadRequest, false, nil, err.Error())
		return
	}
	fmt.Println("Before", limit)
	topCoins, libErr := api.GetFromRatingService(limit)
	if libErr != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, nil, libErr.Error())
		return
	}
	fmt.Println("After")
	response.SendResponse(res, http.StatusOK, true, topCoins, "")
}
