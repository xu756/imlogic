package cache

import (
	"context"
	"encoding/json"
	"imlogic/common/config"
	"time"

	"github.com/redis/go-redis/v9"
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

// Keys 查询 key开头的所有key
func (c *Client) Keys(ctx context.Context, key string) *redis.StringSliceCmd {
	return c.redis.Keys(ctx, key)
}

// DelKeys 删除keys
func (c *Client) DelKeys(ctx context.Context, key ...string) *redis.IntCmd {
	return c.redis.Del(ctx, key...)
}

// Del 删除key
func (c *Client) Del(ctx context.Context, key string) *redis.IntCmd {
	return c.redis.Del(ctx, key)
}
