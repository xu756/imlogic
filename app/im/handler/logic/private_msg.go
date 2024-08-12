package logic

import (
	"context"
	"imlogic/kitex_gen/im"
	"log"
)

// HandlerPrivateMessage implements im.ImHandler.
func (i *ImRpcImpl) HandlerPrivateMessage(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
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
