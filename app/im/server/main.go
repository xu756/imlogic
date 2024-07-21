package main

import (
	"flag"
	"log"
	"net"

	"imlogic/app/im/server/handler"
	"imlogic/app/im/server/rpc"
	"imlogic/common/config"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/im/imserver"

	"github.com/cloudwego/kitex/server"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	// todo 添加处理服务
	go func() {
		handler.NewHub().Run()
	}()

	rpc.Init()

	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.ImServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	svr := imserver.NewServer(handler.NewImServerImpl(),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(middleware.ServerErrorHandler),
	)
	go func() {
		err = svr.Run()
		if err != nil {
			log.Printf(err.Error())
			return
		}
		log.Printf("【 Im-rpc-server 】addr on %s", config.RunData.Addr.ImServerAddr)
	}()
	log.Printf("【 Im-api-server 】addr on %s", config.RunData.Addr.ImAddr)
	handler.InitRouter()
	handler.HttpServer.Spin()

}
