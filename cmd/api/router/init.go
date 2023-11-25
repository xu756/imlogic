package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/xu756/imlogic/common/config"
)

var HttpServer *server.Hertz

func InitRouter() {
	h := server.Default(
		server.WithHostPorts(config.RunData.Addr.ApiAddr),
		server.WithReadBufferSize(1024*1024*100),
		server.WithMaxRequestBodySize(1024*1024*100),
	)

	HttpServer = h
}
