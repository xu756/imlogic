package rpc

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func Init() {
	// FIXME  配置etcd

	//r, err := etcd.NewEtcdResolver(config.RunData.Etcd.Addrs)
	//if err != nil {
	//	hlog.Debug(err)
	//}
	InitPublicClient("public")
}

func ClientErrorHandler(ctx context.Context, err error) error {
	hlog.Debug(err)
	return err
}
