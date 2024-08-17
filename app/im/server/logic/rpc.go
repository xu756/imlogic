package logic

import (
	"context"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/im"
)

// ImServerImpl 服务
type ImServerImpl struct {
}

// SendMsgToGroup implements im.ImServer.
func (i *ImServerImpl) SendMsgToGroup(ctx context.Context, req *im.SendMsgToGroupReq) (res *im.SendMsgToGroupRes, err error) {
	res = &im.SendMsgToGroupRes{}
	for _, linkId := range req.LinkIds {
		ok := service.hub.SendOneMsg(linkId, req.Message)
		if !ok {
			res.LinkIds = append(res.LinkIds, linkId)
		}
	}
	return
}

// SendMsgToOne implements im.ImServer.
func (i *ImServerImpl) SendMsgToOne(ctx context.Context, req *im.SendMsgToOneReq) (res *im.MessageRes, err error) {
	res = &im.MessageRes{}
	ok := service.hub.SendOneMsg(req.LinkId, req.Message)
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
func (i *ImServerImpl) SendMsgToAll(ctx context.Context, req *base.Message) (res *im.MessageRes, err error) {
	service.hub.SendAll(req)
	return &im.MessageRes{Success: true}, nil
}
