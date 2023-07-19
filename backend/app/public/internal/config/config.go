package config

import (
	"github.com/xu756/imlogic/internal/xjwt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Cache     cache.CacheConf
	RedisConf redis.RedisConf
	DbSource  string
	Jwt       xjwt.Jwt
}
