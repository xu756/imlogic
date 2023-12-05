package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/middleware"
	"github.com/xu756/imlogic/internal/result"
)

var HttpServer *server.Hertz

func InitRouter() {
	h := server.Default(
		server.WithHostPorts(config.RunData.Addr.ImAddr),
		server.WithReadBufferSize(1024*1024*100),
		server.WithMaxRequestBodySize(1024*1024*100),
	)
	h.Use(middleware.HertzJwt())
	h.GET("/connect", connect)
	HttpServer = h

}

func connect(ctx context.Context, c *app.RequestContext) {
	err := ConnManager.upgrader.Upgrade(c, ConnManager.handler)
	if err != nil {
		result.HttpError(c, err)
		return
	}

}
