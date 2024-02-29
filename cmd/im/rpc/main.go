package main

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/xu756/imlogic/cmd/im/rpc/handler"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/middleware"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/im/imsrv"
	"log"
	"net"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.ImRpcAddr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := imsrv.NewServer(handler.NewImRpcImpl(),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(xerr.ServerErrorHandler),
		server.WithErrorHandler(middleware.ServerErrorHandler),
	)
	log.Printf("【Im Rpc】 on %s", config.RunData.Addr.ImRpcAddr)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}

}
