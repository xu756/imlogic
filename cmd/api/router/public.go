package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/xu756/imlogic/cmd/api/rpc"
	"github.com/xu756/imlogic/internal/result"
	"github.com/xu756/imlogic/kitex_gen/public"
)

func PublicRoute(r *route.RouterGroup) {
	r.GET("/", Index)
}

func Index(c context.Context, ctx *app.RequestContext) {
	res, err := rpc.PublicClient.LoginByMobile(c, &public.LoginByMobileReq{
		Mobile:    "12345678901",
		Captcha:   "1111",
		SessionId: "111111",
	})
	if err != nil {
		return
	}
	result.HttpSuccess(ctx, res)

}
