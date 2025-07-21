package main

import (
	"consumer/cmd/infra/database"
	"consumer/internal"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	database.InitPostgresConn()
	r := internal.NewMessageLogRepository(database.GetPostgresPool())
	p := internal.NewKafkaLogProcessor(nil, r)

	c, err := internal.NewConsumer(p, &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "create-order",
		"auto.offset.reset": "earliest",
	}, "canonical-orders")

	if err != nil {
		log.Fatal(err)
	}

	if err = c.Consume(); err != nil {
		log.Printf("Error [%s]\n", err.Error())
	}
}
