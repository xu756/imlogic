package handler

import (
	"context"
	"imlogic/common/config"
	"imlogic/internal/xerr"
	"imlogic/kitex_gen/user"
)

func (s *PublicSrvImpl) LoginByPassword(ctx context.Context, req *user.LoginByPasswordReq) (resp *user.LoginRes, err error) {
	userInfo, err := s.Model.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if userInfo.Password != req.Password {
		return nil, xerr.WarnMsg("密码错误")
	}
	token, err := s.Jwt.NewJwtToken(userInfo.ID, userInfo.UUID)
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
	userInfo, err := s.Model.FindUserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	token, err := s.Jwt.NewJwtToken(userInfo.ID, userInfo.UUID)
	if err != nil {
		return nil, err
	}
	resp = &user.LoginRes{
		Token:  token,
		Expire: config.RunData.JwtConfig.Expire,
	}
	return resp, nil
}
