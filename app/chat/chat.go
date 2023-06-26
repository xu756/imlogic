package main

import (
	"flag"
	"fmt"

	"github.com/xu756/imlogic/internal/pb"

	"github.com/xu756/imlogic/app/chat/internal/config"
	"github.com/xu756/imlogic/app/chat/internal/server"
	"github.com/xu756/imlogic/app/chat/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterChatServer(grpcServer, server.NewChatServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	fmt.Printf(" 【chat - 服务】rpc server at %s...\n", c.ListenOn)
	s.Start()
}
