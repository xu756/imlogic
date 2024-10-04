package main

import (
	"log"
	"net"

	"imlogic/app/im/server/handler"
	"imlogic/app/im/server/logic"
	"imlogic/common/config"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/im/imserver"

	"github.com/cloudwego/kitex/server"
)

func main() {

	logic.InitService()

	addr, err := net.ResolveTCPAddr("tcp", config.RunData().Addr.ImServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	svr := imserver.NewServer(
		logic.NewImServerImpl(),
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
	log.Printf("【 Im-api-server 】addr on %s", config.RunData().Addr.ImAddr)
	handler.InitRouter()
	handler.HttpServer.Spin()

}
