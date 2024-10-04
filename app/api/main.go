package main

import (
	"log"

	"imlogic/app/api/logic"
	"imlogic/app/api/router"
	"imlogic/app/api/rpc"
	"imlogic/common/config"
)

func main() {
	router.InitRouter()
	rpc.Init()
	logic.Init()
	log.Printf("【 Api 】addr on %s", config.RunData().Addr.ApiAddr)
	err := router.HttpServer.Run()
	if err != nil {
		log.Print(err)
		return
	}
}
