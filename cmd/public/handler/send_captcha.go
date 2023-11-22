package handler

import (
	"context"
	"github.com/xu756/imlogic/kitex_gen/public"
)

func (s *PublicSrvImpl) GetUserInfo(ctx context.Context, req *public.SendCaptchaReq) (res *public.SendCaptchaRes, err error) {
	return &public.SendCaptchaRes{
		Success: true,
	}, nil
}
