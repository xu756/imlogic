package mongodb

import (
	"context"
	"github.com/xu756/imlogic/common/config"
	"testing"
)

func TestConnect(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	c := NewClient()
	msg, err := c.QueryMsg(context.Background(), "admin", "user-1")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(msg)
}
