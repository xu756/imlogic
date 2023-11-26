package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/xu756/imlogic/common/config"
)

func Init() {
	r, err := etcd.NewEtcdResolver(config.RunData.Etcd.Addrs)
	if err != nil {
		hlog.Debug(err)
	}
	InitPublicClient("public", r)

}
