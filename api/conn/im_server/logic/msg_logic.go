package logic

import (
	"encoding/json"
	"log"

	"github.com/xu756/imlogic/internal/pb"
	"github.com/xu756/imlogic/internal/tool"
	"github.com/xu756/imlogic/internal/xtype"
	"github.com/zeromicro/go-zero/core/logx"
)

func (c *Client) msgLogic(data []byte) {
	var msg xtype.Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		logx.Error("【msgLogic】json.Unmarshal error:" + err.Error())
		return
	}

	//Hubs.Broadcast <- msg
	resp, err := chatRpc.Msg(c.ctx, &pb.Message{
		Uuid:             msg.Uuid,
		ConversationId:   msg.ConversationId,
		ConversationType: pb.ConversationType(msg.ConversationType),
		Sender: &pb.Sender{
			Id:     msg.Sender.Id,
			Name:   msg.Sender.Name,
			Avatar: msg.Sender.Avatar,
			Extra:  msg.Sender.Extra,
		},
		Content:     tool.ToBytes(msg.Content),
		ContentType: pb.MessageContentType(msg.ContentType),
		SendTime:    msg.SendTime,
		Seq:         msg.Seq,
		NeedDecrypt: msg.NeedDecrypt,
		CountUnread: msg.CountUnread,
	})
	if err != nil {
		logx.Error("【msgLogic】ImRpc.Msg error:" + err.Error())
		return
	}
	log.Print(resp)
}
