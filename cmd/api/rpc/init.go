package rpc

import (
	"context"
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
	return err
}
