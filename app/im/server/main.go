package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"imlogic/common/trace"
	"log"
	"net"

	"imlogic/app/im/server/handler"
	"imlogic/app/im/server/logic"
	"imlogic/common/config"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/im/imserver"

	"github.com/cloudwego/kitex/server"
)

const rpcServiceName = "im-server-rpc"
const apiServiceName = "im-server-api"

func main() {

	// rpc
	kitexTracer := trace.KitexTraceSetUp(rpcServiceName)
	defer kitexTracer.Shutdown()
	logic.InitService()
	addr, err := net.ResolveTCPAddr("tcp", config.RunData().Addr.ImServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	svr := imserver.NewServer(
		logic.NewImServerImpl(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: rpcServiceName}),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(middleware.ServerErrorHandler),
	)
	go func() {
		err = svr.Run()
		if err != nil {
			log.Printf(err.Error())
			return
		}
		log.Printf("【 Im-rpc-server 】addr on %s", config.RunData().Addr.ImServerAddr)
	}()

	//api
	hertzTracer := trace.HertzTraceSetUp(apiServiceName)
	defer hertzTracer.Shutdown()
	log.Printf("【 Im-api-server 】addr on %s", config.RunData().Addr.ImAddr)
	handler.InitRouter()
	handler.HttpServer.Spin()

}
