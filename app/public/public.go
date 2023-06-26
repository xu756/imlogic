package main

import (
	"flag"
	"fmt"

	"github.com/xu756/imlogic/app/public/internal/config"
	"github.com/xu756/imlogic/app/public/internal/server"
	"github.com/xu756/imlogic/app/public/internal/svc"
	"github.com/xu756/imlogic/internal/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/public.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterPublicServer(grpcServer, server.NewPublicServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("【public-rpc 服务】 server at %s...\n", c.ListenOn)
	s.Start()
}
