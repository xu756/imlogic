package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
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
	err := ClientManager.upgrader.Upgrade(c, func(ws *websocket.Conn) {
		client := NewClient(ctx, ws, uuid.NewString())
		go client.listenAndWrite()
		ClientManager.register <- client
		client.listenAndRead()
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}

}
