package main

import (
	"flag"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/xu756/imlogic/cmd/im/handler"
	"github.com/xu756/imlogic/common/config"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	// todo 添加处理服务

	handler.NewConnManager("im")
	handler.InitRouter()
	hlog.Debugf("【 Im 】addr on %s", config.RunData.Addr.ImAddr)
	err := handler.HttpServer.Run()
	if err != nil {
		hlog.Debug(err.Error())
		return
	}

}
