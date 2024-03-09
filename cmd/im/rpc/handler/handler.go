package handler

import (
	"github.com/xu756/imlogic/common/cache"
	"github.com/xu756/imlogic/common/db"
	"github.com/xu756/imlogic/common/influxdb"
)

type ImRpcImpl struct {
	Model    db.Model
	Cache    cache.Client
	Influxdb influxdb.Client
}

func NewImRpcImpl() *ImRpcImpl {
	return &ImRpcImpl{
		Model:    db.NewModel(),
		Cache:    cache.NewCacheClient(),
		Influxdb: influxdb.NewClient(),
	}
}
