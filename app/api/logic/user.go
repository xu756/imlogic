package logic

import (
	"context"
	"imlogic/app/api/rpc"
	"imlogic/internal/result"
	"imlogic/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
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
	res, err := rpc.UserClient.GetUserOnlineStatus(ctx, &user.GetOneUserReq{
		Id: req.ID,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res.Status)
}

// 获取当前用户信息
func getUserInfo(ctx context.Context, c *app.RequestContext) {
	userInfo, err := service.Jwt.GetUserInfoFromCookieToken(c)
	if err != nil {
		result.HttpError(c, err)
		return
	}

	res, err := rpc.UserClient.GetOneUserInfo(ctx, &user.GetOneUserReq{
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
	res, err := rpc.UserClient.GetOneUserInfo(ctx, &user.GetOneUserReq{
		Id: req.ID,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)
}
