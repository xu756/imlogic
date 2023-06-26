package config

import (
	"github.com/xu756/imlogic/internal/xjwt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf redis.RedisConf
	Cache     cache.CacheConf
	DbSource  string
	Jwt       xjwt.Jwt
}
