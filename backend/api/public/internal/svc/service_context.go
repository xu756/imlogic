package svc

import (
	"github.com/xu756/imlogic/api/public/internal/config"
	"github.com/xu756/imlogic/app/public/public"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	PublicRpc   public.Public
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		RedisClient: redisClient,
		PublicRpc:   public.NewPublic(zrpc.MustNewClient(c.Public)),
	}
}
