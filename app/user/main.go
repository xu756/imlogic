package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"imlogic/common/trace"
	"log"
	"net"

	"imlogic/app/user/handler"
	"imlogic/common/config"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/user/usersrv"

	"github.com/cloudwego/kitex/server"
)

const rpcServiceName = "im-user"

func main() {
	kitexTracer := trace.KitexTraceSetUp(rpcServiceName)
	defer kitexTracer.Shutdown()
	addr, err := net.ResolveTCPAddr("tcp", config.RunData().Addr.UserAddr)
	if err != nil {
		log.Fatal(err)
	}
	svr := usersrv.NewServer(handler.NewPublicSrvImpl(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: rpcServiceName}),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(middleware.ServerErrorHandler),
	)
	log.Printf("【User Rpc】 on %s", config.RunData().Addr.UserAddr)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
