package main

import (
	"flag"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/xu756/imlogic/cmd/api/router"
	"github.com/xu756/imlogic/cmd/api/rpc"
	"github.com/xu756/imlogic/common/config"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	router.InitRouter()
	rpc.Init()
	hlog.Debugf("【Api】addr on %s", config.RunData.Addr.ApiAddr)
	err := router.HttpServer.Run()
	if err != nil {
		hlog.Debug(err.Error())
		return
	}
}
