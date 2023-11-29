package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/kitex_gen/public/publicsrv"
)

var PublicClient publicsrv.Client

func InitPublicClient(destService string) {
	s, err := publicsrv.NewClient(destService,
		client.WithHostPorts(config.RunData.Addr.PublicAddr),
		client.WithErrorHandler(ClientErrorHandler),
	)
	if err != nil {
		hlog.Debugf("【 连接 public Rpc 失败 】", err)
		panic(err)
	}
	PublicClient = s
}
