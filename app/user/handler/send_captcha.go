package handler

import (
	"context"
	"imlogic/kitex_gen/user"
)

func (s *PublicSrvImpl) SendCaptcha(ctx context.Context, req *user.SendCaptchaReq) (res *user.SendCaptchaRes, err error) {
	return &user.SendCaptchaRes{Success: true}, nil
}
