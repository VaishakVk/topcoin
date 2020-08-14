package main

import (
	"encoding/json"
	"log"
	"time"
	"topcoin/pricingservice/lib/pricelist"
	"topcoin/pricingservice/lib/stocklist"
	"topcoin/pricingservice/messagequeue"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	conn, chann := messagequeue.ConnectAndCreateChannel()
	defer conn.Close()
	defer chann.Close()
	queue := messagequeue.CreateQueue(chann, "stockPrice")

	for {
		stockListMap := stocklist.GetAvailableStocks()
		priceList := pricelist.GetPriceList(stockListMap.Data)
		bytesResult, _ := json.Marshal(priceList.Data)
		payload := messagequeue.GeneratePayload(bytesResult)
		messagequeue.PublishToQueue(chann, queue, payload)
		time.Sleep(time.Second * 5)
	}
}
