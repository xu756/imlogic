package router

import (
	"imlogic/app/api/logic"
	"imlogic/common/config"
	"imlogic/internal/middleware"

	"github.com/cloudwego/hertz/pkg/app/server"
)

var HttpServer *server.Hertz

func InitRouter() {
	h := server.Default(
		server.WithHostPorts(config.RunData.Addr.ApiAddr),
		server.WithReadBufferSize(1024*1024*100),
		server.WithMaxRequestBodySize(1024*1024*100),
	)
	h.Use(middleware.HertzJwt())
	router := h.Group("/api")
	logic.LoginRoute(router.Group("/login"))
	HttpServer = h
}
