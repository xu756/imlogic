package logic

import (
	"context"
	"imlogic/common/types"
	"imlogic/kitex_gen/im"
)

// ImServerImpl 服务
type ImServerImpl struct {
}

func (i ImServerImpl) SendMsg(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	ok := service.hub.SendoneMsg(req.Common.LinkId, types.RpcMsgToMsg(req))
	if !ok {
		return &im.MessageRes{
			Success: false,
		}, nil
	}
	return &im.MessageRes{
		Success: true,
	}, nil
}

// NewImServerImpl 创建服务
func NewImServerImpl() *ImServerImpl {
	return &ImServerImpl{}
}
