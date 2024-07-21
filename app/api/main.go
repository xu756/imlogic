package main

import (
	"flag"
	"log"

	"imlogic/app/api/router"
	"imlogic/app/api/rpc"
	"imlogic/common/config"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	router.InitRouter()
	rpc.Init()
	log.Printf("【 Api 】addr on %s", config.RunData.Addr.ApiAddr)
	err := router.HttpServer.Run()
	if err != nil {
		log.Print(err)
		return
	}
}
