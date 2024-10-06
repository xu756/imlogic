package main

import (
	"imlogic/app/job/logic"
	"imlogic/common/trace"
)

func main() {
	kitexTracer := trace.KitexTraceSetUp("im-server-rpc-client")
	defer kitexTracer.Shutdown()
	logic.InitService()
	logic.StartJob()
}
