package logic

import (
	"context"
	"imlogic/kitex_gen/im"
)

// HandlerGroupMessage implements im.ImHandler.
func (i *ImRpcImpl) HandlerGroupMessage(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	panic("unimplemented")
}
