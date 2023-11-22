package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/xu756/imlogic/cmd/public/handler"
	"github.com/xu756/imlogic/kitex_gen/public/publicsrv"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8681")
	if err != nil {
		klog.Fatal(err)
	}
	svr := publicsrv.NewServer(new(handler.PublicSrvImpl), server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
