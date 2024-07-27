package handler

import (
	"context"
	"imlogic/kitex_gen/im"
)

// TextMessage implements im.ImSrv.
func (i *ImRpcImpl) TextMessage(ctx context.Context, req *im.TextMsg) (res *im.MessageRes, err error) {
	msg := &im.Message{}
	msg.Common = req.Common
	msg.Content = req.Content
	rpc := NewWsServerRpcClient("devLinux")
	go rpc.SendMsgToAll(ctx, msg)
	return &im.MessageRes{}, nil
}
