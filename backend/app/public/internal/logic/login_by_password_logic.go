package logic

import (
	"context"
	"github.com/xu756/imlogic/app/public/internal/svc"
	"github.com/xu756/imlogic/internal/pb"
	"github.com/xu756/imlogic/internal/xerr"
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
	var resp = new(pb.LoginResponse)
	result := l.svcCtx.Captcha.Verify(in.SessionId, in.Code, true)
	if !result {
		return resp, xerr.NewMsgError("验证码错误")
	}
	user, err := l.svcCtx.UserModel.LoginByPassword(l.ctx, in.Username)
	if err != nil {
		return resp, err
	}
	if user.Password != in.Password {
		return resp, xerr.NewMsgError("密码错误")
	}
	roles, err := l.svcCtx.UserRoleModel.FindUserRoles(l.ctx, user.Id)
	if err != nil {
		return nil, err
	}
	jwt, err := l.svcCtx.Jwt.NewJwt(user.Id, roles)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token:  jwt,
		Expire: 7200,
	}, nil
}
