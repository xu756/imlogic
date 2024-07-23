package handler

import (
	"context"
	"imlogic/kitex_gen/im"
)

func (i ImRpcImpl) MetaMessage(ctx context.Context, req *im.MetaMsg) (res *im.MessageRes, err error) {
	res = &im.MessageRes{}
	switch req.Status {
	case im.WsStatus_Connect:
		// todo 连接
	case im.WsStatus_Heartbeat:
		// todo 心跳
	case im.WsStatus_Disconnect:
		// todo 断开连接

	}
	res.Success = true
	res.MsgId = ""
	return
}
