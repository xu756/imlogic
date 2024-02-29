package handler

import (
	"context"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/user"
)

func (s *PublicSrvImpl) LoginByPassword(ctx context.Context, req *user.LoginByPasswordReq) (resp *user.LoginRes, err error) {
	userInfo, err := s.Model.FindUserByUsername(ctx, req.Username, req.Device)
	if err != nil {
		return nil, err
	}
	if userInfo.Password != req.Password {
		return nil, xerr.ErrMsg(xerr.UserPasswordErr)
	}
	token, err := s.Jwt.NewJwtToken(userInfo.UUID, 1.0)
	if err != nil {
		return nil, err
	}
	resp = &user.LoginRes{
		Token:  token,
		Expire: config.RunData.JwtConfig.Expire,
	}
	return resp, nil
}

func (s *PublicSrvImpl) LoginByMobile(ctx context.Context, req *user.LoginByMobileReq) (resp *user.LoginRes, err error) {
	userInfo, err := s.Model.FindUserByMobile(ctx, req.Mobile, req.Device)
	if err != nil {
		return nil, err
	}
	token, err := s.Jwt.NewJwtToken(userInfo.UUID, 1.0)
	if err != nil {
		return nil, err
	}
	resp = &user.LoginRes{
		Token:  token,
		Expire: config.RunData.JwtConfig.Expire,
	}
	return resp, nil
}
