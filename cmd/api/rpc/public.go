package rpc

import "github.com/xu756/imlogic/kitex_gen/public/publicsrv"

var PublicClient publicsrv.Client

func InitPublicClient(destService string) {
	s, err := publicsrv.NewClient(destService)
	if err != nil {
		panic(err)
	}
	PublicClient = s
}
