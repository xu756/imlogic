package logic

import (
	"context"
	"imlogic/kitex_gen/im"
)

// HandlerPrivateMessage implements im.ImHandler.
func (i *ImRpcImpl) HandlerPrivateMessage(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	panic("unimplemented")
}
