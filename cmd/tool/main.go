package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-ping/ping"
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
				go AddrPing(req.Addr, ws)
			}
		})
	})
	h.Spin()
}

func AddrPing(Addr string, ws *websocket.Conn) {
	defer ws.Close()
	pinger := ping.New(Addr)
	pinger.Count = 3
	pinger.OnRecv = func(pkt *ping.Packet) {
		err := ws.WriteJSON(pkt)
		if err != nil {
			return
		}
	}

	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		err := ws.WriteJSON(pkt)
		if err != nil {
			return
		}
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		err := ws.WriteJSON(stats)
		if err != nil {
			return
		}
	}
	pinger.OnSetup = func() {
		err := ws.WriteJSON("ping setup   " + pinger.Addr())
		if err != nil {
			return
		}
	}
	err := pinger.Run()
	if err != nil {
		err := ws.WriteJSON(err.Error())
		if err != nil {
			return
		}
	}

}
