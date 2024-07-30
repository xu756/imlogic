package main

import (
	"flag"
	"log"

	"imlogic/common/config"
	"imlogic/common/mq"

	"github.com/cloudwego/kitex/pkg/klog"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	klog.SetLevel(klog.LevelFatal)
	go Pricate()
	Group()
}

// 私聊消息
func Pricate() {
	rabbitmq, err := mq.NewPrivateMessageMQ()
	defer rabbitmq.Destory()
	if err != nil {
		log.Printf("NewRabbitMQMessage err:%v", err)
	}
	//  rabbitmq.ConsumePrivateChatMessage()
	m, err := rabbitmq.ConsumePrivateMessage()
	if err != nil {
		log.Printf("ConsumePrivateChatMessage err:%v", err)
	}
	for msg := range m {
		log.Printf("%s", msg.Body)
	}
}

// 群聊消息
func Group() {
	rabbitmq, err := mq.NewBroadcastMessageMQ()
	defer rabbitmq.Destory()
	if err != nil {
		log.Printf("NewRabbitMQMessage err:%v", err)
	}
	//  rabbitmq.ConsumePrivateChatMessage()
	m, err := rabbitmq.ConsumeBroadcastMessage()
	if err != nil {
		log.Printf("ConsumePrivateChatMessage err:%v", err)
	}
	for msg := range m {
		log.Printf("%s", msg.Body)
	}
}
