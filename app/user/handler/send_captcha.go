package handler

import (
	"context"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/user"
)

func (s *PublicSrvImpl) SendCaptcha(ctx context.Context, req *user.SendCaptchaReq) (res *base.BoolRes, err error) {
	return &base.BoolRes{Ok: true}, nil
}
