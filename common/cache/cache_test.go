package cache

import (
	"context"
	"encoding/json"
	"github.com/xu756/imlogic/common/config"
	"testing"
)

type cacheUser struct {
	Name string
	Age  int
}

// 序列化
func (m *cacheUser) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// 反序列化
func (m *cacheUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)

}

func TestNewCacheClient(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	client := NewCacheClient()
	ctx := context.Background()
	err := client.Set(ctx, "test", cacheUser{
		Name: "admin",
		Age:  0,
	}, 0)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestNewCacheGet(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	client := NewCacheClient()
	ctx := context.Background()
	var user cacheUser
	err := client.Get(ctx, "test1").Scan(&user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

func TestClient_Keys(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	client := NewCacheClient()
	ctx := context.Background()
	keys := client.Keys(ctx, "conn:*")
	t.Log(keys.Val())
}
