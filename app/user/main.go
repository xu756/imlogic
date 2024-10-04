package main

import (
	"log"
	"net"

	"imlogic/app/user/handler"
	"imlogic/common/config"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/user/usersrv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {

	klog.SetLevel(klog.LevelFatal)
	addr, err := net.ResolveTCPAddr("tcp", config.RunData().Addr.UserAddr)
	if err != nil {
		log.Fatal(err)
	}
	svr := usersrv.NewServer(handler.NewPublicSrvImpl(),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(middleware.ServerErrorHandler),
	)
	log.Printf("【User Rpc】 on %s", config.RunData().Addr.UserAddr)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
