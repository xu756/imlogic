package svc

import (
	"github.com/xu756/imlogic/api/public/internal/config"
	"github.com/xu756/imlogic/app/public/public"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	PublicRpc public.Public
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		PublicRpc: public.NewPublic(zrpc.MustNewClient(c.PublicRpc)),
	}
}
