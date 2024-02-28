package main

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/xu756/imlogic/cmd/user/handler"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/user/usersrv"
	"log"
	"net"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.UserAddr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := usersrv.NewServer(handler.NewPublicSrvImpl(),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(xerr.ServerErrorHandler),
	)
	log.Printf("【User Rpc】 on %s", config.RunData.Addr.UserAddr)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
