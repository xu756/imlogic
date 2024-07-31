package logic

import (
	"context"
	"imlogic/common/cache"
	"imlogic/common/db"
	"imlogic/kitex_gen/im"
)

type ImRpcImpl struct {
	Model db.Model
	Cache *cache.Client
}

// PushMessage implements im.ImSrv.
func (i *ImRpcImpl) PushMessage(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	// rpc := NewWsServerRpcClient("devLinux")
	// go rpc.SendMsgToAll(ctx, req)
	return
}

func NewImRpcImpl() *ImRpcImpl {
	return &ImRpcImpl{
		Model: db.NewModel(),
		Cache: cache.NewClient(),
	}
}
