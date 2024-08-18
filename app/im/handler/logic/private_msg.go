package logic

import (
	"context"
	"github.com/google/uuid"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/im"
	"log"
	"time"
)

// HandlerPrivateMessage implements im.ImHandler.
func (i *ImRpcImpl) HandlerPrivateMessage(ctx context.Context, req *base.Message) (res *im.MessageRes, err error) {
	// 判断是不是好友
	err = i.Model.CheckIsFriend(ctx, req.Sender, req.Receiver)
	if err != nil {
		return &im.MessageRes{
			Success: false,
			Message: &base.Message{
				MsgId:     uuid.NewString(),
				Timestamp: time.Now().Unix(),
				ChatType:  req.ChatType,
				MsgType:   base.MsgType_Event,
				Receiver:  req.Sender,
				Event: &base.Event{
					EventType: base.EventType_NotFriend,
					UserId:    req.Receiver,
				},
			},
		}, nil
	}
	// 保存消息
	err = i.Model.AddOnePrivateMsg(ctx, req.MsgType, req.MsgId, req.Sender, req.Receiver, req.Timestamp, req)
	if err != nil {
		return nil, err
	}
	conns, err := i.Model.GetUserConnsByUserId(ctx, req.Receiver)
	if err != nil {
		return nil, err
	}
	// todo 当不存在连接时，是否需要存储消息
	for _, conn := range conns {
		err = i.PrivateMq.PublishPrivateMessage(conn.LinkID, conn.HostName, req)
		if err != nil {
			log.Print("publish private message error: ", err)
		}
	}
	return &im.MessageRes{
		Success: true,
	}, nil
}
