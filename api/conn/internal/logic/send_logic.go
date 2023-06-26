package logic

import (
	"context"

	"github.com/xu756/imlogic/api/conn/im_server/logic"
	"github.com/xu756/imlogic/internal/pb"
	"github.com/xu756/imlogic/internal/xerr"

	"github.com/xu756/imlogic/api/conn/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Send 其他服务调用发送消息
func (l *SendLogic) Send(in *pb.ConnData) (*pb.ConnResp, error) {
	// todo: add your logic here and delete this line
	client := logic.Hubs.FindOneByImId(in.ConId)
	if client != nil {
		err := client.WriteMessage(in.ConData)
		if err != nil {
			return &pb.ConnResp{
				Ok: false,
			}, xerr.NewMsgError("发送消息失败")
		}
	}
	return &pb.ConnResp{
		Ok: true,
	}, nil
}
