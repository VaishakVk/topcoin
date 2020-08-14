package messagequeue

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

// ConnectAndCreateChannel function
func ConnectAndCreateChannel() (*amqp.Connection, *amqp.Channel) {
	connectionString := os.Getenv("AMQP_URL")
	conn, conErr := amqp.Dial(connectionString)
	if conErr != nil {
		log.Fatal(conErr.Error())
	}

	amqpChannel, chanErr := conn.Channel()
	if chanErr != nil {
		log.Fatal(chanErr.Error())
	}
	return conn, amqpChannel
}

// GeneratePayload - function to generate payload to be sent via Rabbit Channel
func GeneratePayload(data []byte) amqp.Publishing {
	message := amqp.Publishing{
		Body: data,
	}
	return message
}

func CreateQueue(chann *amqp.Channel, name string) amqp.Queue {
	queue, err := chann.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	return queue
}

func PublishToQueue(chann *amqp.Channel, queue amqp.Queue, data amqp.Publishing) {
	err := chann.Publish("", queue.Name, false, false, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Consume(chann *amqp.Channel, queue amqp.Queue) <-chan amqp.Delivery {
	messageChannel, err := chann.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err.Error())
	}
	return messageChannel
}
