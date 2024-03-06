package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/websocket"
)

type pingReq struct {
	Addr string `json:"addr" validate:"required"`
}

var upgrader = websocket.HertzUpgrader{}

func main() {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8080"),
		server.WithReadBufferSize(1024*1024*100),
		server.WithMaxRequestBodySize(1024*1024*100),
	)
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		upgrader.Upgrade(c, func(ws *websocket.Conn) {
			defer ws.Close()
			for {
				var req pingReq
				err := ws.ReadJSON(&req)
				if err != nil {
					return
				}
				go Ping(req.Addr, ws)
			}
		})
	})
	h.Spin()
}
