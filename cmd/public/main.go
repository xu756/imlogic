package main

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/xu756/imlogic/cmd/public/handler"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/kitex_gen/public/publicsrv"
	"log"
	"net"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.PublicAddr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := publicsrv.NewServer(new(handler.PublicSrvImpl),
		server.WithServiceAddr(addr),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
