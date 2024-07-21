package rpc

import (
	"imlogic/common/config"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/user/usersrv"
	"log"

	"github.com/cloudwego/kitex/client"
)

var UserClient usersrv.Client

func InitUserClient(destService string) {
	s, err := usersrv.NewClient(destService,
		client.WithHostPorts(config.RunData.Rpc.UserRpc),
		client.WithErrorHandler(middleware.ClientErrorHandler),
	)

	if err != nil {
		log.Print("【 连接 user Rpc 失败 】", err)
		panic(err)
	}
	UserClient = s
}
