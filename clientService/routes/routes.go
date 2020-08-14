package routes

import (
	"net/http"
	"topcoin/clientservice/controller"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes API
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/rank", controller.GetTopCoins).Methods(http.MethodGet)
}
