package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/xu756/imlogic/cmd/api/rpc"
	"github.com/xu756/imlogic/internal/result"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/public"
)

func PublicRoute(r *route.RouterGroup) {
	r.POST("/login", Login)
}

type loginReq struct {
	LoginType string `json:"loginType,required" vd:"$='account'|$='mobile'"`
	Username  string `json:"userName"`
	Password  string `json:"passWord"`
	Mobile    string `json:"mobile"`
	Captcha   string `json:"captcha"`
	SessionId string `json:"sessionId"`
}

func Login(ctx context.Context, c *app.RequestContext) {

	var req loginReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpError(c, xerr.ParamErr())
		return
	}
	switch req.LoginType {
	case "account":
		if req.Username == "" || req.Password == "" {
			result.HttpError(c, xerr.ParamErr())
			return
		}
		user, err := rpc.PublicClient.LoginByPassword(ctx, &public.LoginByPasswordReq{
			Username:  req.Username,
			Password:  req.Password,
			SessionId: req.SessionId,
		})
		if err != nil {
			result.HttpError(c, err)
			return
		}
		result.HttpSuccess(c, user)
	case "mobile":
		if req.Mobile == "" || req.Captcha == "" || req.SessionId == "" {
			result.HttpError(c, xerr.ParamErr())
			return
		}
		user, err := rpc.PublicClient.LoginByMobile(ctx, &public.LoginByMobileReq{
			Mobile:    req.Mobile,
			Captcha:   req.Captcha,
			SessionId: req.SessionId,
		})
		if err != nil {
			result.HttpError(c, err)
			return
		}
		result.HttpSuccess(c, user)
	}
}
