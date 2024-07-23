package hub

import (
	"imlogic/common/client"
	"imlogic/common/types"
	"os"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
)

type Hub struct {
	clients    sync.Map
	hostName   string
	upgrader   websocket.HertzUpgrader
	broadcast  chan *types.Message
	Register   chan *client.Client
	Unregister chan *client.Client
}

func NewHub() *Hub {
	hostname, _ := os.Hostname()
	hub := &Hub{
		hostName:   hostname,
		broadcast:  make(chan *types.Message, 1024),
		Register:   make(chan *client.Client),
		Unregister: make(chan *client.Client),
		upgrader: websocket.HertzUpgrader{
			CheckOrigin: func(ctx *app.RequestContext) bool {
				return true
			},
		},
	}
	go hub.run()
	return hub
}

func (c *Hub) run() {
	for {
		select {
		case conn := <-c.Register:
			c.add(conn)
		case conn := <-c.Unregister:
			c.del(conn)
		case message := <-c.broadcast:
			c.clients.Range(func(key, value interface{}) bool {
				conn := value.(*client.Client)
				conn.Write(message)
				return true
			})
		}
	}
}

func (h *Hub) UpgradeOneWs(c *app.RequestContext, handler func(ws *websocket.Conn)) error {
	err := h.upgrader.Upgrade(c, handler)
	if err != nil {
		return err
	}
	return nil
}

func (h *Hub) add(conn *client.Client) {
	h.clients.Store(conn.LinkId, conn)
}

func (h *Hub) del(conn *client.Client) {
	h.clients.Delete(conn.LinkId)
}

func (h *Hub) SendoneMsg(LinkId string, msg *types.Message) bool {
	conn, ok := h.clients.Load(LinkId)
	if ok {
		conn.(*client.Client).Send <- msg
		return true
	}
	return false
}
