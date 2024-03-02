package handler

import (
	"context"
	"github.com/xu756/imlogic/common/types"
	"github.com/xu756/imlogic/kitex_gen/im"
)

// ImServerImpl 服务
type ImServerImpl struct {
}

func (i ImServerImpl) Send(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	ok, conn := ClientManager.GetConn(req.GetTo())
	if ok {
		conn.writer <- &types.Message{
			MsgId:     req.MsgId,
			Device:    req.Device,
			Timestamp: req.Timestamp,
			Params:    req.Params,
			Action:    req.Action,
			From:      req.From,
			To:        req.To,
			MsgType:   req.MsgType,
			MsgMeta: types.MsgMeta{
				DetailType: req.MsgMeta.DetailType,
				Version:    req.MsgMeta.Version,
				Interval:   req.MsgMeta.Interval,
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
		Success: ok,
	}, nil
}

// NewImServerImpl 创建服务
func NewImServerImpl() *ImServerImpl {
	return &ImServerImpl{}
}
