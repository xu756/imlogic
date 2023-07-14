package logic

import (
	"context"
	"github.com/xu756/imlogic/internal/pb"

	"github.com/xu756/imlogic/app/public/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过密码登录
func (l *LoginByPasswordLogic) LoginByPassword(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginResponse{}, nil
}
