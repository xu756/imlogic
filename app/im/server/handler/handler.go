package handler

import (
	"imlogic/app/im/server/logic"
	"imlogic/common/config"

	"github.com/cloudwego/hertz/pkg/app/server"
)

var HttpServer *server.Hertz

// InitRouter 路由
func InitRouter() {
	h := server.Default(
		server.WithHostPorts(config.RunData().Addr.ImAddr),
		server.WithReadBufferSize(1024*1024*100),
		server.WithMaxRequestBodySize(1024*1024*100),
	)
	router := h.Group("/api/im")
	router.GET("/connect", logic.Connect)
	HttpServer = h
}
