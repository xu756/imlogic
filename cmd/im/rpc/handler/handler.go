package handler

import (
	"github.com/xu756/imlogic/common/cache"
	"github.com/xu756/imlogic/common/db"
)

type ImRpcImpl struct {
	Model db.Model
	Cache cache.Client
}

func NewImRpcImpl() *ImRpcImpl {
	return &ImRpcImpl{
		Model: db.NewModel(),
		Cache: cache.NewCacheClient(),
	}
}
