package handler

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"imlogic/app/im/server/logic"
	"imlogic/common/config"
)

var HttpServer *server.Hertz

// InitRouter 路由
func InitRouter() {
	tracer, cfg := tracing.NewServerTracer()
	h := server.Default(tracer,
		server.WithHostPorts(config.RunData().Addr.ImAddr),
		server.WithReadBufferSize(1024*1024*100),
		server.WithMaxRequestBodySize(1024*1024*100),
	)
	h.Use(tracing.ServerMiddleware(cfg))
	router := h.Group("/api/ws")
	router.GET("/connect", logic.Connect)
	HttpServer = h
}
