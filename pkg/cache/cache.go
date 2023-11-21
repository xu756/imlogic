package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var cdb *cache.Cache

func InitCache() {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	cdb = c

}

type UserCache struct {
	Email     string
	Code      string
	SessionId string
}

// Get 获取缓存
func Get(key string) (*UserCache, bool) {
	v, ok := cdb.Get(key)
	if !ok {
		return &UserCache{}, false
	}
	return v.(*UserCache), true
}

// Set 设置缓存
func Set(key, email, code, sessionId string) {
	cdb.Set(key, &UserCache{
		Email:     email,
		Code:      code,
		SessionId: sessionId,
	}, 60*time.Second)

}
