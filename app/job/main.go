package main

import (
	"flag"
	"log"

	"imlogic/common/config"
	"imlogic/common/mq"
	"imlogic/kitex_gen/im"

	"github.com/cloudwego/kitex/pkg/klog"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	klog.SetLevel(klog.LevelFatal)
	go Pricate()
	Broadcast()
	// PubSub()
	// Delay()
}

func privateFunc(linkId, HostName string, msg *im.Message) {
	log.Printf("linkId:%s, HostName:%s, msg:%s", linkId, HostName, msg.Content)
}

// 私聊消息
func Pricate() {
	rabbitmq, err := mq.NewPrivateMessageMQ()
	defer rabbitmq.Destory()
	if err != nil {
		log.Printf("NewRabbitMQMessage err:%v", err)
	}
	//  rabbitmq.ConsumePrivateChatMessage()
	err = rabbitmq.ConsumePrivateMessage(privateFunc)
	if err != nil {
		log.Printf("ConsumePrivateChatMessage err:%v", err)
	}
}
func broadcastFunc(msg *im.Message) {
	log.Printf("msg:%s", msg.Content)
}

// 广播消息
func Broadcast() {
	rabbitmq, err := mq.NewBroadcastMessageMQ()
	defer rabbitmq.Destory()
	if err != nil {
		log.Printf("NewRabbitMQMessage err:%v", err)
	}
	//  rabbitmq.ConsumePrivateChatMessage()
	err = rabbitmq.ConsumeBroadcastMessage(broadcastFunc)
	if err != nil {
		log.Printf("ConsumePrivateChatMessage err:%v", err)
	}
}

// 延迟队列
func Delay() {
	rabbitmq, err := mq.NewDelayMessageMQ("delay")
	defer rabbitmq.Destory()
	if err != nil {
		log.Printf("NewRabbitMQMessage err:%v", err)
	}
	m, err := rabbitmq.ConsumeDelayMessage()
	if err != nil {
		log.Printf("ConsumeDelayMessage err:%v", err)
	}
	for msg := range m {
		log.Printf("%s", msg.Body)
	}
}
