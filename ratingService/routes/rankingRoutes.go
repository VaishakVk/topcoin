package routes

import (
	"net/http"
	"topcoin/ratingservice/controller"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	rankingRoutes := router.PathPrefix("/ranking").Subrouter()
	rankingRoutes.HandleFunc("", controller.GetTopCoins).Methods(http.MethodGet)
}
