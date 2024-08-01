package logic

import (
	"context"
	"imlogic/common/cache"
	"imlogic/common/db"
	"imlogic/common/mq"
	"imlogic/kitex_gen/im"
	"log"
)

type ImRpcImpl struct {
	Model db.Model
	Cache *cache.Client
	Mq    *mq.RabbitMQ
}

// PushMessage implements im.ImSrv.
func (i *ImRpcImpl) PushMessage(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	err = i.Mq.PublishBroadcastMessage(req)
	if err != nil {
		log.Printf("PushMessage err:%v", err)
	}
	return &im.MessageRes{
		Success: true,
	}, nil
}

func NewImRpcImpl() *ImRpcImpl {
	rabbitmq, err := mq.NewBroadcastMessageMQ()
	if err != nil {
		log.Printf("NewRabbitMQMessage err:%v", err)
	}
	return &ImRpcImpl{
		Model: db.NewModel(),
		Cache: cache.NewClient(),
		Mq:    rabbitmq,
	}
}
