package main

import (
	"github.com/go-ping/ping"
	"github.com/hertz-contrib/websocket"
)

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
