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
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	var priceList map[string]model.Price

	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	errorsHandledRouter := middlewares.Recovery(router)
	headersAddedRouter := middlewares.SetHeaders(errorsHandledRouter)

	conn, chann := messagequeue.ConnectAndCreateChannel()
	defer conn.Close()
	defer chann.Close()

	queue := messagequeue.CreateQueue(chann, "stockPrice")
	messageChannel := messagequeue.Consume(chann, queue)

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

	http.ListenAndServe(":8000", headersAddedRouter)
}
