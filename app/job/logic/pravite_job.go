package logic

import (
	"context"
	"imlogic/common/mq"
	"imlogic/kitex_gen/im"
	"log"
)

// 私聊消息
func private() {
	rabbitmq, err := mq.NewPrivateMessageMQ()
	defer rabbitmq.Destory()
	if err != nil {
		log.Printf("NewRabbitMQMessage err:%v", err)
	}
	err = rabbitmq.ConsumePrivateMessage(privateFunc)
	if err != nil {
		log.Printf("ConsumePrivateChatMessage err:%v", err)
	}
}

func privateFunc(linkId, hostName string, msg *im.Message) {
	rpc := NewWsServerRpcClient("hostName")
	res, err := rpc.SendMsgToOne(context.Background(), &im.SendMsgTooneReq{
		LinkId:  linkId,
		Message: msg,
	})
	if err != nil {
		log.Printf("SendMsgToOne err:%v", err)
	}
	if res.Success {
		log.Printf("SendMsgToOne success")
	} else {
		log.Printf("SendMsgToOne fail")
	}
}
