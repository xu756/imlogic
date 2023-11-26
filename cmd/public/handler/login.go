package handler

import (
	"context"
	"github.com/xu756/imlogic/internal/tool"
	"github.com/xu756/imlogic/kitex_gen/public"
)

func (s *PublicSrvImpl) LoginByPassword(ctx context.Context, req *public.LoginByPasswordReq) (resp *public.LoginRes, err error) {

	return resp, nil
}

func (s *PublicSrvImpl) LoginByMobile(ctx context.Context, req *public.LoginByMobileReq) (res *public.LoginRes, err error) {
	res = &public.LoginRes{
		Token:  tool.TimeNowString(),
		Expire: 6400,
	}
	return res, nil
}
