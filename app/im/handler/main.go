package main

import (
	"flag"
	"log"
	"net"

	"imlogic/app/im/handler/logic"
	"imlogic/common/config"
	"imlogic/internal/middleware"
	"imlogic/kitex_gen/im/imhandler"

	"github.com/cloudwego/kitex/server"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.ImHandlerAddr)
	if err != nil {
		log.Fatal(err)

	}
	svr := imhandler.NewServer(logic.NewImRpcImpl(),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(middleware.ServerErrorHandler),
	)
	log.Printf("【Im Rpc】 on %s", config.RunData.Addr.ImHandlerAddr)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}

}
