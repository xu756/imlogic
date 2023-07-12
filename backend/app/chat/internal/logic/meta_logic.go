package logic

import (
	"context"

	"github.com/xu756/imlogic/internal/pb"

	"github.com/xu756/imlogic/app/chat/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type MetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MetaLogic {
	return &MetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 元事件 连接 断开 状态更新 解密错误
func (l *MetaLogic) Meta(in *pb.ImMeta) (*pb.MetaResp, error) {
	// todo: add your logic here and delete this line

	return &pb.MetaResp{}, nil
}
