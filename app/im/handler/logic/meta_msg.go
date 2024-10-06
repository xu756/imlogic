package logic

import (
	"context"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/im"
)

func (i ImRpcImpl) MetaMessage(ctx context.Context, req *base.MetaMsg) (res *im.MessageRes, err error) {
	res = &im.MessageRes{}
	switch req.Status {
	case base.WsStatus_Connect:
		err = i.Model.AddUserConn(ctx, req.UserId, req.HostIp, req.Device, req.LinkId)
		if err != nil {
			return nil, err
		}
		err = i.Model.DeleteUserConnByHeartbeatTime(ctx)
		if err != nil {
			return nil, err
		}
	case base.WsStatus_Heartbeat:
		err = i.Model.UpdateLastHeartbeatTime(ctx, req.UserId, req.HostIp, req.Device, req.LinkId)
		if err != nil {
			return nil, err
		}
	case base.WsStatus_Disconnect:
		err = i.Model.DeleteUserConn(ctx, req.UserId, req.HostIp, req.LinkId)
		if err != nil {
			return nil, err
		}
	}
	res.Success = true
	return
}
