package mq

import (
	"imlogic/common/config"
	"log"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestNewClient(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	log.Print(config.RunData.MqUrl)
	client := NewClient("im/message")
	go func() {
		err := client.Consume(func(delivery amqp.Delivery) {
			log.Printf("Received a message: %s\n", delivery.Body)
		})
		if err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(5 * time.Second) // Wait for consumer to be ready.
	err := client.Publish("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second) // Wait for consumer to be ready.
	err = client.Publish("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(15 * time.Second) // Wait for messages to be processed.

}
