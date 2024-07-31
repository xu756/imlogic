package logic

import (
	"imlogic/common/cache"
	"imlogic/common/db"
)

var service *Service

type Service struct {
	Model db.Model
	Cache *cache.Client
}

func InitService() {
	service = &Service{
		Model: db.NewModel(),
		Cache: cache.NewClient(),
	}
}
