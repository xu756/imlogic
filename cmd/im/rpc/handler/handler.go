package handler

import (
	"github.com/xu756/imlogic/common/cache"
	"github.com/xu756/imlogic/common/db"
	"github.com/xu756/imlogic/common/mongodb"
)

type ImRpcImpl struct {
	Model   db.Model
	Cache   cache.Client
	mongodb mongodb.Client
}

func NewImRpcImpl() *ImRpcImpl {
	return &ImRpcImpl{
		Model:   db.NewModel(),
		Cache:   cache.NewCacheClient(),
		mongodb: mongodb.NewClient(),
	}
}
