package handler

import (
	"imlogic/common/cache"
	"imlogic/common/db"
	"imlogic/common/mongodb"
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
