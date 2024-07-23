package logic

import (
	"context"
	"imlogic/common/client"
	"imlogic/common/types"
	"imlogic/kitex_gen/im"
	"log"
)

func Msglogic(c *client.Client, msg *types.Message) {
	if msg.MsgMeta.Detail == "heartbeat" {
		// todo 心跳
		return
	}
	commonMsg := &im.CommonMsg{
		MsgId:    msg.MsgId,
		LinkId:   msg.LinkId,
		ChatType: im.ChatType(msg.ChatType),
		Sender:   msg.Sender,
		ChatId:   msg.ChatId,
		GroupId:  msg.GroupId,
		MsgType:  im.MsgType(msg.MsgType),
	}
	switch msg.MsgType {
	case types.Text:
		var textMsg = new(im.TextMsg)
		textMsg.Common = commonMsg
		textMsg.Content = msg.MsgContent.Content
		_, err := service.ImSrv.TextMessage(c.Ctx, textMsg)
		if err != nil {
			log.Print("send text message failed", err)
			return
		}

	case types.Image:
		var imgMsg = new(im.ImgMsg)
		imgMsg.Common = commonMsg
		var imgs []*im.ImageType
		for _, img := range msg.MsgContent.Img {
			imgs = append(imgs, &im.ImageType{
				Uid: img.UID,
				Url: img.URL,
			})
		}
		imgMsg.Img = imgs
		_, err := service.ImSrv.ImageMessage(c.Ctx, imgMsg)
		if err != nil {
			log.Print("send image message failed", err)
			return
		}
	case types.File:
		// todo 文件消息
		var fileMsg = new(im.FileMsg)
		fileMsg.Common = commonMsg
		fileMsg.File = &im.FileType{
			Uid: msg.MsgContent.File.UID,
			Url: msg.MsgContent.File.URL,
		}
		_, err := service.ImSrv.FileMessage(c.Ctx, fileMsg)
		if err != nil {
			log.Print("send file message failed", err)
			return
		}
	case types.Audio:
		// todo 音频消息
		var audioMsg = new(im.AudioMsg)
		audioMsg.Common = commonMsg
		audioMsg.Audio = &im.AudioType{
			Uid:      msg.MsgContent.Audio.UID,
			Url:      msg.MsgContent.Audio.URL,
			Duration: msg.MsgContent.Audio.Duration,
		}
		_, err := service.ImSrv.AudioMessage(c.Ctx, audioMsg)
		if err != nil {
			log.Print("send audio message failed", err)
			return
		}
	case types.Video:
		// todo 视频消息
		var videoMsg = new(im.VideoMsg)
		videoMsg.Common = commonMsg
		videoMsg.Video = &im.VideoType{
			Uid:      msg.MsgContent.Video.UID,
			Url:      msg.MsgContent.Video.URL,
			Duration: msg.MsgContent.Video.Duration,
		}
		_, err := service.ImSrv.VideoMessage(c.Ctx, videoMsg)
		if err != nil {
			log.Print("send video message failed", err)
			return
		}

	}
}

// rpcMetaMsg
func MetaMsg(ctx context.Context, c *client.Client, msg *im.MetaMsg) {
	res, err := service.ImSrv.MetaMessage(ctx, msg)
	if err != nil || res.Success == false {
		log.Print("send meta message failed", err)
		c.Ctx.Done()
		return
	}

}
