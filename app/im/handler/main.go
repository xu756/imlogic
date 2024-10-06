package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"imlogic/common/trace"
	"log"
	"net"

	"imlogic/app/im/handler/logic"
	"imlogic/common/config"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/im/imhandler"

	"github.com/cloudwego/kitex/server"
)

const rpcServiceName = "im-handler"

func main() {
	kitexTracer := trace.KitexTraceSetUp(rpcServiceName)
	defer kitexTracer.Shutdown()

	addr, err := net.ResolveTCPAddr("tcp", config.RunData().Addr.ImHandlerAddr)
	if err != nil {
		log.Fatal(err)

	}
	svr := imhandler.NewServer(logic.NewImRpcImpl(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: rpcServiceName}),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(middleware.ServerErrorHandler),
	)
	log.Printf("【Im Rpc】 on %s", config.RunData().Addr.ImHandlerAddr)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}

}
