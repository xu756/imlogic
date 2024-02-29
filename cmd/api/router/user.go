package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/xu756/imlogic/cmd/api/rpc"
	"github.com/xu756/imlogic/internal/result"
	"github.com/xu756/imlogic/kitex_gen/user"
)

func UserRoute(r *route.RouterGroup) {
	r.POST("/account", loginByPassword)
	r.POST("/mobile", loginByMobile)
	r.POST("/captcha", sendCaptcha)
}

func loginByPassword(ctx context.Context, c *app.RequestContext) {
	var req LoginReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	if req.Username == "" || req.Password == "" {
		result.HttpParamErr(c)
		return
	}
	res, err := rpc.UserClient.LoginByPassword(ctx, &user.LoginByPasswordReq{
		Username:  req.Username,
		Password:  req.Password,
		SessionId: req.SessionId,
		Device:    req.Device,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)
}

func loginByMobile(ctx context.Context, c *app.RequestContext) {
	var req LoginReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	if req.Mobile == "" || req.Captcha == "" {
		result.HttpParamErr(c)
		return
	}
	res, err := rpc.UserClient.LoginByMobile(ctx, &user.LoginByMobileReq{
		Mobile:    req.Mobile,
		Captcha:   req.Captcha,
		SessionId: req.SessionId,
		Device:    req.Device,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)

}

func sendCaptcha(ctx context.Context, c *app.RequestContext) {
	var req SendCaptchaReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	res, err := rpc.UserClient.SendCaptcha(ctx, &user.SendCaptchaReq{
		Mobile:    req.Mobile,
		SessionId: req.SessionId,
	})
	if err != nil {
		result.HttpParamErr(c)
		return
	}
	result.HttpSuccess(c, res.Success)
}
