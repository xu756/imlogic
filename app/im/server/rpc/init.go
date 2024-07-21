package rpc

func Init() {
	// FIXME  配置etcd

	//r, err := etcd.NewEtcdResolver(config.RunData.Etcd.Addrs)
	//if err != nil {
	//	hlog.Debug(err)
	//}
	InitImSrvClient("im-server")
}
