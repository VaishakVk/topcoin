package controller

import (
	"net/http"
	"strconv"
	"topcoin/ratingservice/ranking"
	"topcoin/ratingservice/response"
)

func GetTopCoins(res http.ResponseWriter, req *http.Request) {
	keys, exist := req.URL.Query()["limit"]
	if !exist {
		response.SendResponse(res, http.StatusBadRequest, false, "Limit is invalid")
		return
	}

	limit, err := strconv.ParseInt(keys[0], 10, 32)
	if err != nil {
		response.SendResponse(res, http.StatusBadRequest, false, err.Error())
		return
	}
	if limit == 0 {
		response.SendResponse(res, http.StatusBadRequest, false, "Limit is invalid")
		return
	}
	topCoins, libErr := ranking.GetTopCoins(limit)
	if libErr != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, libErr.Error())
		return
	}
	response.SendResponse(res, http.StatusOK, true, topCoins)
}
