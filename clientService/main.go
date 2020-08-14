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
	errorsHandledRouter := middlewares.Recovery(router)
	urlRouter := middlewares.RemoveTrailingSlashes(errorsHandledRouter)
	headersAddedRouter := middlewares.SetHeaders(urlRouter)
	log.Fatal(http.ListenAndServe(":8001", headersAddedRouter))
}
