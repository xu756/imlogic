package logic

import (
	"imlogic/common/config"
	"imlogic/common/hub"
	"imlogic/internal/middleware"
	"imlogic/internal/xjwt"
	"imlogic/kitex_gen/im/imsrv"
	"log"

	"github.com/cloudwego/kitex/client"
)

var service *Service

type Service struct {
	ImSrv imsrv.Client
	Jwt   *xjwt.Jwt
	hub   *hub.Hub
}

func InitService() {
	s, err := imsrv.NewClient("im-server",
		client.WithHostPorts(config.RunData.Rpc.ImRpcAddr),
		client.WithErrorHandler(middleware.ClientErrorHandler),
	)
	if err != nil {
		log.Print("【 连接 im-server Rpc 失败 】", err)
		panic(err)
	}
	service = &Service{
		ImSrv: s,
		Jwt:   xjwt.NewJwt(),
		hub:   hub.NewHub(),
	}
}
