package controller

import (
	"net/http"
	"strconv"
	"topcoin/ratingservice/ranking"
	"topcoin/ratingservice/response"
)

// GetTopCoins - API Handler to get top coins based on limit parameter
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
	if limit == 0 {
		response.SendResponse(res, http.StatusBadRequest, false, nil, "Limit is invalid")
		return
	}
	topCoins, libErr := ranking.GetTopCoins(limit)
	if libErr != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, nil, libErr.Error())
		return
	}
	response.SendResponse(res, http.StatusOK, true, topCoins, "")
}
