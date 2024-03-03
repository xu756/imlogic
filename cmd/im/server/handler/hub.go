package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"github.com/xu756/imlogic/common/types"
	"os"
	"sync"
)

var Hub = newHub()

func newHub() *hub {
	hostname, _ := os.Hostname()
	return &hub{
		HostName:   hostname,
		broadcast:  make(chan *types.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		upgrader: websocket.HertzUpgrader{
			CheckOrigin: func(ctx *app.RequestContext) bool {
				return true
			},
		},
	}
}

func (c *hub) Run() {
	for {
		select {
		case conn := <-c.register:
			c.add(conn)
		case conn := <-c.unregister:
			c.del(conn)
		case message := <-c.broadcast:
			message.Action = "broadcast"
			c.Clients.Range(func(key, value interface{}) bool {
				conn := value.(*Client)
				message.To = "all"
				conn.Write(message)
				return true
			})
		}
	}
}

type hub struct {
	Clients    sync.Map
	HostName   string
	upgrader   websocket.HertzUpgrader
	broadcast  chan *types.Message
	register   chan *Client
	unregister chan *Client
}

func (c *hub) add(conn *Client) {
	c.Clients.Store(conn.linkID, conn)
}

func (c *hub) del(conn *Client) {
	c.Clients.Delete(conn.linkID)
}

func (c *hub) GetConn(linkID string) (bool, *Client) {
	conn, ok := c.Clients.Load(linkID)
	if ok {
		return true, conn.(*Client)
	}
	return false, nil
}
