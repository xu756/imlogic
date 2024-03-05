package handler

import (
	"context"
	"github.com/xu756/imlogic/kitex_gen/im"
)

func (i ImRpcImpl) Receive(ctx context.Context, req *im.Message) (res *im.Message, err error) {
	res = &im.Message{
		MsgId:     req.MsgId,
		From:      req.From,
		To:        req.To,
		Timestamp: req.Timestamp,
		Params: map[string]string{
			"send": "ok",
		},
	}
	return
}
