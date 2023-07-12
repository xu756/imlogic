package logic

import (
	"context"

	"github.com/xu756/imlogic/internal/xerr"

	"github.com/xu756/imlogic/app/public/internal/svc"
	"github.com/xu756/imlogic/internal/pb"

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

func (l *LoginByPasswordLogic) LoginByPassword(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var resp = new(pb.LoginResponse)
	resp.Expire = l.svcCtx.Config.Jwt.Expire
	user, err := l.svcCtx.UserModel.LoginByPassword(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	//todo:密码加密
	if user.Password != in.Password {
		return nil, xerr.NewMsgError("密码错误")
	}
	resp.Token, err = l.svcCtx.Jwt.NewJwt(user.Id, user.Role)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
