package handler

import (
	"context"
	"github.com/xu756/imlogic/kitex_gen/public"
)

func (s *PublicSrvImpl) GetUserInfo(ctx context.Context, req *public.UserInfoReq) (res *public.UserInfoRes, err error) {
	return &public.UserInfoRes{
		Username: "xu756",
	}, nil
}
