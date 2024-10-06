package logic

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/im/imserver"
	"log"

	"github.com/cloudwego/kitex/client"
)

func StartJob() {
	go private()
	broadcast()
}

func NewWsServerRpcClient(hostIp string) imserver.Client {
	s, err := imserver.NewClient("im-server",
		client.WithSuite(tracing.NewClientSuite()),
		client.WithHostPorts(hostIp+":7083"),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "im-server-rpc-client"}),
		client.WithErrorHandler(middleware.ClientErrorHandler),
	)
	if err != nil {
		log.Print("【 连接 im-server Rpc 失败 】", err)
		panic(err)
	}
	return s
}
