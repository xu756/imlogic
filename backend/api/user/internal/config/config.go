package config

import (
	"github.com/xu756/imlogic/internal/xjwt"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	User zrpc.RpcClientConf
	Jwt  xjwt.Jwt
}
