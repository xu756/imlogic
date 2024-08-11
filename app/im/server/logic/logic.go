package logic

import (
	"context"
	"imlogic/common/client"
	"imlogic/common/types"
	"imlogic/kitex_gen/im"
	"log"
)

func Msglogic(c *client.Client, msg *types.Message) {
	if msg.ChatType == types.SystemMessage && msg.Content == "ping" {
		// todo 心跳
		c.ResetHeartbeat()
		MetaMsg(c, c.HeartbeatMsg())
		c.SendMsg(&types.Message{
			ChatType: types.SystemMessage,
			Content:  "pong",
		})
		return
	}
	ctx := context.Background()
	medias := make([]*im.MediaType, 0)
	for _, media := range msg.Media {
		medias = append(medias, &im.MediaType{
			Uid: media.UID,
			Url: media.URL,
		})
	}
	var res = &im.MessageRes{}
	var err error
	switch msg.ChatType {
	case types.PrivateChat:
		res, err = service.ImHandler.HandlerPrivateMessage(ctx, &im.Message{
			LinkId:    msg.LinkId,
			MsgId:     msg.MsgId,
			Timestamp: msg.Timestamp,
			ChatType:  im.ChatType(msg.ChatType),
			Sender:    msg.Sender,
			Receiver:  msg.Receiver,
			ChatId:    msg.ChatId,
			GroupId:   msg.GroupId,
			Content:   msg.Content,
			MsgType:   im.MsgType(msg.MsgType),
			Media:     medias,
		})

	case types.GroupChat:
		res, err = service.ImHandler.HandlerGroupMessage(ctx, &im.Message{
			LinkId:    msg.LinkId,
			MsgId:     msg.MsgId,
			Timestamp: msg.Timestamp,
			ChatType:  im.ChatType(msg.ChatType),
			Sender:    msg.Sender,
			Receiver:  msg.Receiver,
			ChatId:    msg.ChatId,
			GroupId:   msg.GroupId,
			Content:   msg.Content,
			MsgType:   im.MsgType(msg.MsgType),
			Media:     medias,
		})
	}
	if err != nil {
		log.Print("message failed", err)
		return
	}
	if !res.Success {
		log.Print("message failed")
		return
	}
}

// rpcMetaMsg
func MetaMsg(c *client.Client, msg *im.MetaMsg) {
	ctx := context.Background()
	res, err := service.ImHandler.MetaMessage(ctx, msg)
	if err != nil {
		log.Print("send meta message failed", err)
		ctx.Done()
		return
	}
	if !res.Success {
		log.Print("send meta message failed", res)
		ctx.Done()
		return
	}

}
