package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"topcoin/ratingservice/messagequeue"
	"topcoin/ratingservice/model"

	"github.com/joho/godotenv"
)

func main() {
	var priceList map[string]model.Price
	var priceRank []model.PriceRanking

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	conn, chann := messagequeue.ConnectAndCreateChannel()
	defer conn.Close()
	defer chann.Close()

	queue := messagequeue.CreateQueue(chann, "stockPrice")
	messageChannel := messagequeue.Consume(chann, queue)

	// stockRank := messagequeue.CreateQueue(chann, "stockRank")

	// go func() {
	for response := range messageChannel {
		// fmt.Println("Herer", string(response.Body))
		parseErr := json.Unmarshal(response.Body, &priceList)
		if parseErr != nil {
			log.Fatal(parseErr.Error())
		}

		if err := response.Ack(false); err != nil {
			log.Printf("Error acknowledging message : %s", err)
		} else {
			log.Printf("Acknowledged message")
		}

		for _, value := range priceList {
			priceRank = append(priceRank, model.PriceRanking{Name: value.Name, Symbol: value.Symbol, PriceInUSD: value.Quote.USD.Price})
		}

		sort.Slice(priceRank, func(i, j int) bool {
			return priceRank[i].PriceInUSD > priceRank[j].PriceInUSD
		})
		fmt.Println("Soreted Data", priceRank)
	}
	// }()
}
