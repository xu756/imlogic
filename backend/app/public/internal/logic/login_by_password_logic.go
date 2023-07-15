package logic

import (
	"context"
	"github.com/xu756/imlogic/internal/pb"
	"github.com/xu756/imlogic/internal/xerr"

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

// LoginByPassword 通过密码登录
func (l *LoginByPasswordLogic) LoginByPassword(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	result := l.svcCtx.Captcha.Verify(in.SessionId, in.Code, true)
	if !result {
		return nil, xerr.NewMsgError("验证码错误")
	}
	return &pb.LoginResponse{
		Token:  "token",
		Expire: 7200,
	}, nil
}
