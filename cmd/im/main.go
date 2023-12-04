package main

import (
	"flag"
	"github.com/xu756/imlogic/cmd/im/handler"
	"github.com/xu756/imlogic/common/config"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	// todo 添加处理服务
	handler.NewConnManager("im")

}
