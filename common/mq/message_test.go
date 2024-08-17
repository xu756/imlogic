package mq

import (
	"fmt"
	"imlogic/common/config"
	"imlogic/kitex_gen/base"
	"testing"
)

func TestPrivateChatMessage(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	rabbitmq, err := NewPrivateMessageMQ()
	defer rabbitmq.Destory()
	if err != nil {
		t.Errorf("NewRabbitMQMessage err:%v", err)
	}
	msg := new(base.Message)
	for i := 0; i < 3; i++ {
		msg.Content = fmt.Sprintf("private msg: %d", i)
		err = rabbitmq.PublishPrivateMessage("user_001", "devLinux", msg)
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
	msg := new(base.Message)
	for i := 0; i < 10; i++ {
		msg.Content = fmt.Sprintf("broadcast msg: %d", i)
		err = rabbitmq.PublishBroadcastMessage(msg)
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

func TestNewDelayMessageMQ(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	rabbitmq, err := NewDelayMessageMQ("delay")
	defer rabbitmq.Destory()
	if err != nil {
		t.Errorf("NewDelayMessageMQ err:%v", err)
	}
	err = rabbitmq.PublishDelayMessage(3, "delay msg")
	if err != nil {
		t.Errorf("PublishDelayMessage err:%v", err)
	}

}
