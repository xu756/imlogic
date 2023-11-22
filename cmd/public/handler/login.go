package handler

import (
	"context"
	"github.com/xu756/imlogic/kitex_gen/public"
)

func (s *PublicSrvImpl) LoginByPassword(ctx context.Context, req *public.LoginByPasswordReq) (resp *public.LoginRes, err error) {
	resp = &public.LoginRes{
		Token:  req.Username,
		Expire: 6400,
	}
	return resp, nil
}

func (s *PublicSrvImpl) LoginByEmail(ctx context.Context, req *public.LoginByEmailReq) (res *public.LoginRes, err error) {
	res = &public.LoginRes{
		Token:  req.Email,
		Expire: 6400,
	}
	return res, nil
}
