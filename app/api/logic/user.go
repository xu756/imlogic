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
	// r.POST("/list", getUserList)
	// r.POST("/add", addUser)
	// r.POST("/edit", editUser)
	// r.POST("/delete", deleteUser)
	// r.POST("/reset", resetPassword)
	r.POST("/info", getUserInfo)
}

// 获取单个用户信息
func getUserInfo(ctx context.Context, c *app.RequestContext) {
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
