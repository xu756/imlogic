package logic

import (
	"context"
	"imlogic/common/types"
	"imlogic/kitex_gen/im"
)

// ImServerImpl 服务
type ImServerImpl struct {
}

// SendMsgToGroup implements im.ImServer.
func (i *ImServerImpl) SendMsgToGroup(ctx context.Context, req *im.SendMsgToGroupReq) (res *im.SendMsgToGroupRes, err error) {
	res = &im.SendMsgToGroupRes{}
	for _, linkId := range req.LinkIds {
		ok := service.hub.SendoneMsg(linkId, types.RpcMsgToMsg(req.Message))
		if !ok {
			res.LinkIds = append(res.LinkIds, linkId)
		}
	}
	return
}

// SendMsgToOne implements im.ImServer.
func (i *ImServerImpl) SendMsgToOne(ctx context.Context, req *im.SendMsgTooneReq) (res *im.MessageRes, err error) {
	res = &im.MessageRes{}
	ok := service.hub.SendoneMsg(req.LinkId, types.RpcMsgToMsg(req.Message))
	if !ok {
		res.Success = false
		return
	}
	res.Success = true
	return
}

// NewImServerImpl 创建服务
func NewImServerImpl() *ImServerImpl {
	return &ImServerImpl{}
}

// SendMsgToAll implements im.ImServer.
func (i *ImServerImpl) SendMsgToAll(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	service.hub.SendAll(types.RpcMsgToMsg(req))
	return &im.MessageRes{Success: true}, nil
}
