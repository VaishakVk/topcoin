package main

import (
	"encoding/json"
	"log"
	"net/http"
	"topcoin/ratingservice/messagequeue"
	"topcoin/ratingservice/middlewares"
	"topcoin/ratingservice/model"
	"topcoin/ratingservice/ranking"
	"topcoin/ratingservice/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	var priceList map[string]model.Price

	// HTTP setup
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	errorsHandledRouter := middlewares.Recovery(router)
	headersAddedRouter := middlewares.SetHeaders(errorsHandledRouter)

	// Message queue setup
	conn, chann := messagequeue.ConnectAndCreateChannel()
	defer conn.Close()
	defer chann.Close()

	queue := messagequeue.CreateQueue(chann, "stockPrice")
	messageChannel := messagequeue.Consume(chann, queue)

	// Trigger a go routine to listen to price data
	go func() {
		for response := range messageChannel {
			parseErr := json.Unmarshal(response.Body, &priceList)
			if parseErr != nil {
				log.Fatal(parseErr.Error())
			}

			if err := response.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

			ranking.GenerateRanking(priceList)
		}
	}()

	// Listen to HTTP requests
	log.Fatal(http.ListenAndServe(":8000", headersAddedRouter))
}
