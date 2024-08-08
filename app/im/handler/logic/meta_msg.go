package logic

import (
	"context"
	"imlogic/kitex_gen/im"
)

func (i ImRpcImpl) MetaMessage(ctx context.Context, req *im.MetaMsg) (res *im.MessageRes, err error) {
	res = &im.MessageRes{}
	switch req.Status {
	case im.WsStatus_Connect:
		err = i.Model.AddUserConn(ctx, req.UserId, req.HostName, req.Device, req.LinkId)
		if err != nil {
			return nil, err
		}
		err = i.Model.DeleteUserConnByHeartbeatTime(ctx)
		if err != nil {
			return nil, err
		}
	case im.WsStatus_Heartbeat:
		err = i.Model.UpdateLastHeartbeatTime(ctx, req.UserId, req.HostName, req.Device, req.LinkId)
		if err != nil {
			return nil, err
		}
	case im.WsStatus_Disconnect:
		err = i.Model.DeleteUserConn(ctx, req.UserId, req.HostName, req.Device, req.LinkId)
		if err != nil {
			return nil, err
		}
	}

	res.Success = true
	res.MsgId = ""
	return
}
