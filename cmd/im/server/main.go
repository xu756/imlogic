package main

import (
	"flag"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/xu756/imlogic/cmd/im/server/handler"
	"github.com/xu756/imlogic/cmd/im/server/rpc"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/middleware"
	"github.com/xu756/imlogic/kitex_gen/im/imserver"
	"log"
	"net"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	// todo 添加处理服务
	go handler.NewHub().Run()
	handler.InitRouter()
	rpc.Init()

	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.ImServerAddr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := imserver.NewServer(handler.NewImServerImpl(),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(middleware.ServerErrorHandler),
	)
	hlog.Infof("【 Im-api-server 】addr on %s", config.RunData.Addr.ImAddr)

	hlog.Infof("【 Im-rpc-server 】addr on %s", config.RunData.Addr.ImServerAddr)

	go func() {
		err = svr.Run()
		if err != nil {
			log.Printf(err.Error())
			return
		}
	}()
	handler.HttpServer.Spin()
}
