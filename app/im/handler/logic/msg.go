package logic

import (
	"imlogic/common/cache"
	"imlogic/common/db"
	"imlogic/common/mq"
	"log"
)

type ImRpcImpl struct {
	Model       db.Model
	Cache       *cache.Client
	BroadcastMq *mq.RabbitMQ
	PrivateMq   *mq.RabbitMQ
}

func NewImRpcImpl() *ImRpcImpl {
	broadcastRabbitmq, err := mq.NewBroadcastMessageMQ()
	if err != nil {
		log.Printf("NewBroadcastMessageMQ err:%v", err)
		return nil
	}
	privateRabbitmq, err := mq.NewPrivateMessageMQ()
	if err != nil {
		log.Printf("NewPrivateMessageMQ err:%v", err)
		return nil
	}
	return &ImRpcImpl{
		Model:       db.NewModel(),
		Cache:       cache.NewClient(),
		BroadcastMq: broadcastRabbitmq,
		PrivateMq:   privateRabbitmq,
	}
}
