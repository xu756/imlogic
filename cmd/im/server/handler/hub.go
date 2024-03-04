package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"github.com/xu756/imlogic/common/types"
	"os"
	"sync"
)

var hub *Hub

func NewHub() *Hub {
	hostname, _ := os.Hostname()
	hub = &Hub{
		HostName:   hostname,
		broadcast:  make(chan *types.Message, 1024),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		upgrader: websocket.HertzUpgrader{
			CheckOrigin: func(ctx *app.RequestContext) bool {
				return true
			},
		},
	}
	return hub
}

func (c *Hub) Run() {
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
				conn.send <- message
				return true
			})
		}
	}
}

type Hub struct {
	Clients    sync.Map
	HostName   string
	upgrader   websocket.HertzUpgrader
	broadcast  chan *types.Message
	register   chan *Client
	unregister chan *Client
}

func (c *Hub) add(conn *Client) {
	c.Clients.Store(conn.linkID, conn)
}

func (c *Hub) del(conn *Client) {

	c.Clients.Delete(conn.linkID)
}

func (c *Hub) GetConn(linkID string) (bool, *Client) {
	conn, ok := c.Clients.Load(linkID)
	if ok {
		return true, conn.(*Client)
	}
	return false, nil
}
