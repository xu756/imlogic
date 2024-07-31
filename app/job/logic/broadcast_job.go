package logic

import (
	"imlogic/common/mq"
	"imlogic/kitex_gen/im"
	"log"
)

// 广播消息
func broadcast() {
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

func broadcastFunc(msg *im.Message) {
	log.Printf("msg:%s", msg.Content)
}
