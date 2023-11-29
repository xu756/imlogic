package handler

import (
	"context"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/public"
)

func (s *PublicSrvImpl) LoginByPassword(ctx context.Context, req *public.LoginByPasswordReq) (resp *public.LoginRes, err error) {
	user, err := s.Model.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != req.Password {
		return nil, xerr.ErrMsg(xerr.UserPasswordErr)
	}
	token, err := s.Jwt.NewJwtToken(user.ID, []uint64{1})
	if err != nil {
		return nil, err
	}
	resp = &public.LoginRes{
		Token:  token,
		Expire: config.RunData.JwtConfig.Expire,
	}
	return resp, nil
}

func (s *PublicSrvImpl) LoginByMobile(ctx context.Context, req *public.LoginByMobileReq) (res *public.LoginRes, err error) {

	return res, nil
}
