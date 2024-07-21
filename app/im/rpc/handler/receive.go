package handler

import (
	"context"
	"imlogic/kitex_gen/im"
)

func (i ImRpcImpl) Receive(ctx context.Context, req *im.Message) (res *im.Message, err error) {
	err = i.mongodb.CreateMsg(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &im.Message{
		LinkId:    req.LinkId,
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
