package main

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/xu756/imlogic/cmd/im/rpc/handler"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/im/imsrv"
	"log"
	"net"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	// FIXME: 添加etcd注册
	//r, err := etcd.NewEtcdRegistry(config.RunData.Etcd.Addrs)
	//if err != nil {
	//	log.Fatal(err)
	//}
	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.ImRpcAddr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := imsrv.NewServer(handler.NewImRpcImpl(),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(xerr.ServerErrorHandler),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "im-rpc",
		}),
		//server.WithRegistry(r),
	)
	log.Printf("【Im Rpc】 on %s", config.RunData.Addr.ImRpcAddr)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}

}
