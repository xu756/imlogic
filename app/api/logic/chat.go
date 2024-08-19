package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"imlogic/app/api/rpc"
	"imlogic/internal/result"
	"imlogic/kitex_gen/base"
)

func ChatRoute(r *route.RouterGroup) {
	r.POST("/list", getChatList)

}

// 获取聊天列表
func getChatList(ctx context.Context, c *app.RequestContext) {
	userInfo, err := service.Jwt.GetUserInfoFromHeardToken(c)
	if err != nil {
		result.HttpError(c, err)
		return
	}
	res, err := rpc.UserClient.GetUserChatList(ctx, &base.GetOneReq{Id: userInfo.UserId})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)
}
