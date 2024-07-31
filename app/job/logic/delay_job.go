package logic

import (
	"imlogic/common/mq"
	"log"
)

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
