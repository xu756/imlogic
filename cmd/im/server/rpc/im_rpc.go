package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/im/imsrv"
)

var ImSrvClient imsrv.Client

func InitImSrvClient(destService string) {
	s, err := imsrv.NewClient(destService,
		client.WithHostPorts(config.RunData.Rpc.ImRpcAddr),
		client.WithErrorHandler(xerr.ClientErrorHandler),
	)
	if err != nil {
		hlog.Debug("【 连接 im-server Rpc 失败 】", err)
		panic(err)
	}
	ImSrvClient = s
}
