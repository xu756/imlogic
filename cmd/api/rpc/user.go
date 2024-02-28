package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/user/usersrv"
)

var UserClient usersrv.Client

func InitUserClient(destService string) {
	s, err := usersrv.NewClient(destService,
		client.WithHostPorts(config.RunData.Rpc.UserRpc),
		client.WithErrorHandler(xerr.ClientErrorHandler),
	)

	if err != nil {
		hlog.Debug("【 连接 user Rpc 失败 】", err)
		panic(err)
	}
	UserClient = s
}
