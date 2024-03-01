package cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/xu756/imlogic/common/config"
	"time"
)

type Client struct {
	redis  *redis.Client
	Prefix string
}

func NewCacheClient() Client {
	return Client{
		redis: redis.NewClient(&redis.Options{
			Addr:     config.RunData.RedisConfig.Addr,
			Password: config.RunData.RedisConfig.Password,
			DB:       config.RunData.RedisConfig.Db,
		}),
		Prefix: config.RunData.RedisConfig.Prefix,
	}
}

func (c *Client) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.redis.Get(ctx, key)
}
func (c *Client) Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	value, _ := json.Marshal(v)
	return c.redis.Set(ctx, key, value, expiration).Err()
}
