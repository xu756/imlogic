package logic

import (
	"context"

	"github.com/xu756/imlogic/internal/pb"

	"github.com/xu756/imlogic/app/chat/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type MsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgLogic {
	return &MsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// im消息
func (l *MsgLogic) Msg(in *pb.Message) (*pb.MsgResp, error) {
	// todo: add your logic here and delete this line

	return &pb.MsgResp{}, nil
}
