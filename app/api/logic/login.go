package logic

import (
	"context"

	"imlogic/app/api/rpc"
	"imlogic/internal/result"
	"imlogic/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func LoginRoute(r *route.RouterGroup) {
	r.POST("/account", loginByPassword)
	r.POST("/mobile", loginByMobile)
	r.POST("/captcha", sendCaptcha)
}

func loginByPassword(ctx context.Context, c *app.RequestContext) {
	var req user.LoginByPasswordReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	if req.Username == "" || req.Password == "" {
		result.HttpParamErr(c)
		return
	}
	res, err := rpc.UserClient.LoginByPassword(ctx, &req)
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)
}

func loginByMobile(ctx context.Context, c *app.RequestContext) {
	var req user.LoginByMobileReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	if req.Mobile == "" || req.Captcha == "" {
		result.HttpParamErr(c)
		return
	}
	res, err := rpc.UserClient.LoginByMobile(ctx, &req)
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)

}

func sendCaptcha(ctx context.Context, c *app.RequestContext) {
	var req user.SendCaptchaReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	res, err := rpc.UserClient.SendCaptcha(ctx, &req)
	if err != nil {
		result.HttpParamErr(c)
		return
	}
	result.HttpSuccess(c, res.Ok)
}
