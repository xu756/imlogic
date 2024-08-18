package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"imlogic/app/api/rpc"
	"imlogic/internal/result"
	"imlogic/kitex_gen/base"
)

func UserRouter(r *route.RouterGroup) {
	r.POST("/oneInfo", getUserOneInfo)
	r.POST("/info", getUserInfo)
	r.POST("/status", getUserStatus)
}

// 获取用户在线状态
func getUserStatus(ctx context.Context, c *app.RequestContext) {
	var req GetOneUser
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	res, err := rpc.UserClient.GetUserOnlineStatus(ctx, &base.GetOneReq{
		Id: req.ID,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res.Ok)
}

// 获取当前用户信息
func getUserInfo(ctx context.Context, c *app.RequestContext) {
	userInfo, err := service.Jwt.GetUserInfoFromHeardToken(c)
	if err != nil {
		result.HttpError(c, err)
		return
	}

	res, err := rpc.UserClient.GetOneUserInfo(ctx, &base.GetOneReq{
		Id: userInfo.UserId,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)
}

// 获取单个用户信息
func getUserOneInfo(ctx context.Context, c *app.RequestContext) {
	var req GetOneUser
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	res, err := rpc.UserClient.GetOneUserInfo(ctx, &base.GetOneReq{
		Id: req.ID,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)
}
