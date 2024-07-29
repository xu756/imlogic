package main

import (
	"flag"
	"log"

	"imlogic/common/config"
	"imlogic/common/mq"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/cloudwego/kitex/pkg/klog"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	klog.SetLevel(klog.LevelFatal)
	client := mq.NewClient("im/message")

	for {
		err := client.Consume(func(delivery amqp.Delivery) {
			log.Printf("Received 1: %s\n", delivery.Body)
		})
		if err != nil {
			log.Print(err)
		}
	}
}
