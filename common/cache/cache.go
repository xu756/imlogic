package cache

import (
	"imlogic/common/config"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	redis  *redis.Client
	prefix string
}

func NewClient() *Client {
	return &Client{
		redis: redis.NewClient(&redis.Options{
			Addr:     config.RunData.RedisConfig.Addr,
			Password: config.RunData.RedisConfig.Password,
			DB:       config.RunData.RedisConfig.Db,
		}),
		prefix: config.RunData.RedisConfig.Prefix,
	}
}
