package main

import (
	"imlogic/common/trace"
	"log"

	"imlogic/app/api/logic"
	"imlogic/app/api/router"
	"imlogic/app/api/rpc"
	"imlogic/common/config"
)

const serviceName = "api"

func main() {
	router.InitRouter()
	rpc.Init()
	logic.Init()
	hertzTracer := trace.HertzTraceSetUp(serviceName)
	defer hertzTracer.Shutdown()

	log.Printf("【 Api 】addr on %s", config.RunData().Addr.ApiAddr)
	router.HttpServer.Spin()
}
