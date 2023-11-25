package main

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xu756/imlogic/cmd/public/handler"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/kitex_gen/public/publicsrv"
	"log"
	"net"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	r, err := etcd.NewEtcdRegistry(config.RunData.Etcd.Addrs)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.PublicAddr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := publicsrv.NewServer(new(handler.PublicSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "public",
			Addr:        addr,
		}),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
