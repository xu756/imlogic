package logic

import (
	"context"
	"github.com/xu756/imlogic/internal/pb"

	"github.com/xu756/imlogic/app/user/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.Empty) (*pb.UserInfo, error) {

	return &pb.UserInfo{}, nil
}
