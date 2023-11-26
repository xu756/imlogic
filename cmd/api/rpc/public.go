package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/xu756/imlogic/kitex_gen/public/publicsrv"
)

var PublicClient publicsrv.Client

func InitPublicClient(destService string, r discovery.Resolver) {
	s := publicsrv.MustNewClient(destService, client.WithResolver(r))
	PublicClient = s
}
