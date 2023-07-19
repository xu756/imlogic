package main

import (
	"flag"
	"fmt"

	"github.com/xu756/imlogic/api/user/internal/config"
	"github.com/xu756/imlogic/api/user/internal/handler"
	"github.com/xu756/imlogic/api/user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("【 用户 api 启动 】 at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
