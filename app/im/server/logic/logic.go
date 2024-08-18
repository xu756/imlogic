package logic

import (
	"context"
	"imlogic/common/client"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/im"
	"log"
)

func MsgLogic(c *client.Client, msg *base.Message) {
	if msg.ChatType == base.ChatType_SystemMessage && msg.Content == "ping" {
		// todo 心跳
		c.ResetHeartbeat()
		MetaMsg(c, c.HeartbeatMsg())
		c.SendMsg(&base.Message{
			ChatType: base.ChatType_SystemMessage,
			Content:  "pong",
		})
		return
	}
	ctx := context.Background()
	medias := make([]*base.MediaType, 0)
	for _, media := range msg.Media {
		medias = append(medias, &base.MediaType{
			Uid: media.Uid,
			Url: media.Url,
		})
	}
	var res = &im.MessageRes{}
	var err error
	switch msg.ChatType {
	case base.ChatType_PrivateChat:
		res, err = service.ImHandler.HandlerPrivateMessage(ctx, &base.Message{
			MsgId:     msg.MsgId,
			Timestamp: msg.Timestamp,
			ChatType:  msg.ChatType,
			Sender:    msg.Sender,
			Receiver:  msg.Receiver,
			Content:   msg.Content,
			MsgType:   msg.MsgType,
			Media:     medias,
		})

	case base.ChatType_GroupChat:
		res, err = service.ImHandler.HandlerGroupMessage(ctx, &base.Message{
			MsgId:     msg.MsgId,
			Timestamp: msg.Timestamp,
			ChatType:  msg.ChatType,
			Sender:    msg.Sender,
			Receiver:  msg.Receiver,
			Content:   msg.Content,
			MsgType:   msg.MsgType,
			Media:     medias,
		})
	}
	if err != nil {
		log.Print("message failed", err)
		return
	}
	if !res.Success {
		c.SendMsg(res.Message)
		return
	}
}

// rpcMetaMsg
func MetaMsg(c *client.Client, msg *base.MetaMsg) {
	ctx := context.Background()
	res, err := service.ImHandler.MetaMessage(ctx, msg)
	if err != nil {
		log.Print("send meta message failed", err)
		ctx.Done()
		return
	}
	if !res.Success {
		log.Print("send meta message failed", res)
		c.SendMsg(res.Message)
		ctx.Done()
		return
	}

}
