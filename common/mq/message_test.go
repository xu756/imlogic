package mq

import (
	"fmt"
	"imlogic/common/config"
	"testing"
)

func TestPrivateChatMessage(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	rabbitmq, err := NewPrivateMessageMQ()
	defer rabbitmq.Destory()
	if err != nil {
		t.Errorf("NewRabbitMQMessage err:%v", err)
	}
	for i := 0; i < 100; i++ {
		err = rabbitmq.PublishPrivateMessage(fmt.Sprintf("private msg: %d", i))
		if err != nil {
			t.Errorf("PubPrivateChatMessage err:%v", err)
		}
	}
}

func TestBroadcastMessage(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	rabbitmq, err := NewBroadcastMessageMQ()
	defer rabbitmq.Destory()
	if err != nil {
		t.Errorf("NewRabbitMQMessage err:%v", err)
	}
	for i := 0; i < 100; i++ {
		err = rabbitmq.PublishBroadcastMessage(fmt.Sprintf("broadcast msg: %d", i))
		if err != nil {
			t.Errorf("PubBroadcastMessage err:%v", err)
		}
	}
}

func TestNewRabbitMQPubSub(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	rabbitmq := NewRabbitMQPubSub("im")
	defer rabbitmq.Destory()
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub(fmt.Sprintf("pubsub msg: %d", i))
	}
}
