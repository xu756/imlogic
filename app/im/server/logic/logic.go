package logic

import (
	"context"
	"imlogic/common/client"
	"imlogic/common/types"
	"imlogic/kitex_gen/im"
	"log"
)

func Msglogic(c *client.Client, msg *types.Message) {
	if msg.ChatType != types.SystemMessage && msg.Content == "heartbeat" {
		// todo 心跳
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
	_, err := service.ImHandler.PushMessage(ctx, &im.Message{
		LinkId:    msg.LinkId,
		MsgId:     msg.MsgId,
		Timestamp: msg.Timestamp,
		ChatType:  im.ChatType(msg.ChatType),
		Sender:    msg.Sender,
		ChatId:    msg.ChatId,
		GroupId:   msg.GroupId,
		Content:   msg.Content,
		MsgType:   im.MsgType(msg.MsgType),
		Media:     medias,
	})
	if err != nil {
		log.Print("send text message failed", err)
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
