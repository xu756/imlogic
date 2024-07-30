package mq

import (
	"fmt"
	"imlogic/common/config"
	"testing"
)

func TestSimple(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	rabbitmq := NewRabbitMQSimple("im")
	defer rabbitmq.Destory()
	for i := 0; i < 10; i++ {
		rabbitmq.PublishSimple(fmt.Sprintf("Hello World! %d", i))
	}
}

func TestPubSub(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	rabbitmq := NewRabbitMQPubSub("pubsub")
	defer rabbitmq.Destory()
	for i := 0; i < 10; i++ {
		rabbitmq.PublishPub("Hello World! " + fmt.Sprintf("%d", i))
	}
}
