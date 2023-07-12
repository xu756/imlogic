package svc

import (
	"github.com/xu756/imlogic/api/conn/internal/config"
	"github.com/xu756/imlogic/app/chat/chat"
	"github.com/xu756/imlogic/internal/xjwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	Jwt         xjwt.Jwt
	Chat        chat.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		RedisClient: redisClient,
		Config:      c,
		Jwt:         c.Jwt,
		Chat:        chat.NewChat(zrpc.MustNewClient(c.Chat)),
	}
}
