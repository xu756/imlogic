package logic

import (
	"context"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/im"
	"log"
)

// HandlerPrivateMessage implements im.ImHandler.
func (i *ImRpcImpl) HandlerPrivateMessage(ctx context.Context, req *base.Message) (res *im.MessageRes, err error) {
	// 判断是不是好友
	err = i.Model.CheckIsFriend(ctx, req.ChatId, req.Sender, req.Receiver)
	if err != nil {
		return nil, err
	}
	// 保存消息
	err = i.Model.AddOnePrivateMsg(ctx, int32(req.MsgType), req.MsgId, req.ChatId, req.Sender, req.Timestamp, req)
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
