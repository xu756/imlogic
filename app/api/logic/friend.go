package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"imlogic/app/api/rpc"
	"imlogic/internal/result"
	"imlogic/kitex_gen/base"
	"imlogic/kitex_gen/user"
)

func FriendRoute(r *route.RouterGroup) {
	r.POST("/list", getFriendList)
	r.POST("/add", addFriend)

}

// 获取好友列表
func getFriendList(ctx context.Context, c *app.RequestContext) {
	userInfo, err := service.Jwt.GetUserInfoFromHeardToken(c)
	if err != nil {
		result.HttpError(c, err)
		return
	}
	var req user.GetFriendListReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	req.Id = userInfo.UserId
	res, err := rpc.UserClient.GetUserFriendList(ctx, &req)
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)
}

// 添加好友
func addFriend(ctx context.Context, c *app.RequestContext) {
	userInfo, err := service.Jwt.GetUserInfoFromHeardToken(c)
	if err != nil {
		result.HttpError(c, err)
		return
	}
	var req base.GetOneReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpParamErr(c)
		return
	}
	res, err := rpc.UserClient.AddFriend(ctx, &user.AddFriendReq{
		Owner:  userInfo.UserId,
		WithId: req.Id,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res.Ok)
}
