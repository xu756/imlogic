package hub

import (
	"imlogic/common/client"
	"imlogic/kitex_gen/base"
	"os"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
)

type Hub struct {
	clients sync.Map
	// clients    map[string]*client.Client
	lock       sync.Mutex
	hostName   string
	upgrader   websocket.HertzUpgrader
	broadcast  chan *base.Message
	register   chan *client.Client
	unregister chan *client.Client
}

func NewHub() *Hub {
	hostname, _ := os.Hostname()
	hub := &Hub{
		hostName:   hostname,
		broadcast:  make(chan *base.Message, 1024),
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
			h.clients.Store(conn.GetLinkId(), conn)
		case conn := <-h.unregister:
			h.clients.Delete(conn.GetLinkId())
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

func (h *Hub) SendOneMsg(LinkId string, msg *base.Message) bool {
	conn, ok := h.clients.Load(LinkId)
	if ok {
		conn.(*client.Client).SendMsg(msg)
		return true
	}
	return false
}

func (h *Hub) SendAll(msg *base.Message) {
	h.lock.Lock()
	h.clients.Range(func(key, value interface{}) bool {
		c := value.(*client.Client)
		c.SendMsg(msg)
		return true
	})
	h.lock.Unlock()
}
