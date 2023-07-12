package config

import (
	"github.com/xu756/imlogic/internal/xjwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type IMConfig struct {
	Port      int `json:"port"`
	Heartbeat int `json:"heartbeat"`
}

type Config struct {
	zrpc.RpcServerConf
	ImConfig  IMConfig
	RedisConf redis.RedisConf
	Jwt       xjwt.Jwt
	Chat      zrpc.RpcClientConf
}
