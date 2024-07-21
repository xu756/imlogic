package rpc

import (
	"imlogic/common/config"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/im/imsrv"
	"log"

	"github.com/cloudwego/kitex/client"
)

var ImSrvClient imsrv.Client

func InitImSrvClient(destService string) {
	s, err := imsrv.NewClient(destService,
		client.WithHostPorts(config.RunData.Rpc.ImRpcAddr),
		client.WithErrorHandler(middleware.ClientErrorHandler),
	)
	if err != nil {
		log.Print("【 连接 im-server Rpc 失败 】", err)
		panic(err)
	}
	ImSrvClient = s
}
