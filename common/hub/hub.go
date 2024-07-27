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
	register   chan *client.Client
	unregister chan *client.Client
	sync.RWMutex
}

func NewHub() *Hub {
	hostname, _ := os.Hostname()
	hub := &Hub{
		hostName:   hostname,
		broadcast:  make(chan *types.Message, 1024),
		register:   make(chan *client.Client),
		unregister: make(chan *client.Client),
		upgrader: websocket.HertzUpgrader{
			CheckOrigin: func(ctx *app.RequestContext) bool {
				return true
			},
		},
	}
	go hub.run()
	return hub
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.register:
			h.clients.Store(conn.LinkId, conn)
		case conn := <-h.unregister:
			h.clients.Delete(conn.LinkId)
		case message := <-h.broadcast:
			h.clients.Range(func(key, value interface{}) bool {
				conn := value.(*client.Client)
				conn.SendMsg(message)
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

func (h *Hub) AddOneClient(conn *client.Client) {
	h.register <- conn

}

func (h *Hub) DelOneClient(conn *client.Client) {
	h.unregister <- conn
}

func (h *Hub) SendoneMsg(LinkId string, msg *types.Message) bool {
	conn, ok := h.clients.Load(LinkId)
	if ok {
		conn.(*client.Client).SendMsg(msg)
		return true
	}
	return false
}

func (h *Hub) SendAll(msg *types.Message) {
	h.clients.Range(func(key, value interface{}) bool {
		value.(*client.Client).SendMsg(msg)
		return true
	})
}
