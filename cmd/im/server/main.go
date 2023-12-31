package main

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/xu756/imlogic/cmd/im/server/handler"
	"github.com/xu756/imlogic/cmd/im/server/rpc"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/im/imserver"
	"log"
	"net"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	// todo 添加处理服务
	handler.NewClientnManager()
	handler.InitRouter()
	rpc.Init()

	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.ImServerAddr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := imserver.NewServer(handler.NewImServerImpl(),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(xerr.ServerErrorHandler),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "im-server",
		}),
		//server.WithRegistry(r),
	)
	log.Printf("【 Im-api-server 】addr on %s", config.RunData.Addr.ImAddr)

	go handler.HttpServer.Spin()
	log.Printf("【 Im-rpc-server 】 on %s", config.RunData.Addr.ImServerAddr)

	err = svr.Run()
	if err != nil {
		log.Printf(err.Error())
		return
	}
}
