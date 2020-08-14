package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"topcoin/clientservice/middlewares"
	"topcoin/clientservice/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load Env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Http server setup
	router := mux.NewRouter()

	routes.RegisterUserRoutes(router)
	errorsHandled := middlewares.Recovery(router)
	urlUpdated := middlewares.RemoveTrailingSlashes(errorsHandled)
	headersAdded := middlewares.SetHeaders(urlUpdated)
	log.Fatal(http.ListenAndServe(":6667", headersAdded))
}
