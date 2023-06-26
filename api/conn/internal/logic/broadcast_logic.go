package logic

import (
	"context"

	"github.com/xu756/imlogic/api/conn/im_server/logic"
	"github.com/xu756/imlogic/internal/pb"

	"github.com/xu756/imlogic/api/conn/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type BroadcastLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBroadcastLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BroadcastLogic {
	return &BroadcastLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Broadcast 其他服务调用广播消息
func (l *BroadcastLogic) Broadcast(in *pb.ConnData) (*pb.ConnResp, error) {
	logic.Hubs.Broadcast <- in.ConData
	return &pb.ConnResp{
		Ok: true,
	}, nil
}
