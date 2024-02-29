package handler

import (
	"context"
	"github.com/xu756/imlogic/kitex_gen/user"
)

func (s *PublicSrvImpl) SendCaptcha(ctx context.Context, req *user.SendCaptchaReq) (res *user.SendCaptchaRes, err error) {
	return &user.SendCaptchaRes{Success: true}, nil
}
