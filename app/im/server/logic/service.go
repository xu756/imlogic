package logic

import (
	"imlogic/common/config"
	"imlogic/common/hub"
	"imlogic/internal/middleware"
	"imlogic/internal/xjwt"
	"imlogic/kitex_gen/im/imhandler"
	"log"

	"github.com/cloudwego/kitex/client"
)

var service *Service

type Service struct {
	ImHandler imhandler.Client
	Jwt       *xjwt.Jwt
	hub       *hub.Hub
}

func InitService() {
	s, err := imhandler.NewClient("im-handler",
		client.WithHostPorts(config.RunData().Rpc.ImHandlerAddr),
		client.WithErrorHandler(middleware.ClientErrorHandler),
	)
	if err != nil {
		log.Print("【 连接 im Rpc 失败 】", err)
		panic(err)
	}
	service = &Service{
		ImHandler: s,
		Jwt:       xjwt.NewJwt(),
		hub:       hub.NewHub(),
	}
}
