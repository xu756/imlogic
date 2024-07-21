package handler

import (
	"context"
	"imlogic/common/types"
	"imlogic/kitex_gen/im"
)

// ImServerImpl 服务
type ImServerImpl struct {
}

func (i ImServerImpl) Send(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	ok, conn := hub.GetConn(req.GetTo())
	if ok {
		conn.send <- &types.Message{
			MsgId:     req.MsgId,
			Timestamp: req.Timestamp,
			Params:    req.Params,
			Action:    req.Action,
			LinkId:    conn.linkID,
			From:      req.From,
			To:        req.To,
			MsgType:   req.MsgType,
			MsgMeta: types.MsgMeta{
				DetailType: req.MsgMeta.DetailType,
				Version:    req.MsgMeta.Version,
			},
			MsgContent: types.MsgContent{
				DetailType: req.MsgContent.DetailType,
				Text:       req.MsgContent.Text,
				ImgUrl:     req.MsgContent.ImgUrl,
				AudioUrl:   req.MsgContent.AudioUrl,
				VideoUrl:   req.MsgContent.VideoUrl,
			},
		}
	}
	return &im.MessageRes{
		Success: true,
	}, nil
}

// NewImServerImpl 创建服务
func NewImServerImpl() *ImServerImpl {
	return &ImServerImpl{}
}
