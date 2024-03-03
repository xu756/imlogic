package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"github.com/xu756/imlogic/common/types"
	"os"
	"sync"
)

var ClientManager = new(clientManager)

func NewClientManager() {
	hostname, _ := os.Hostname()
	ClientManager = &clientManager{
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
	go ClientManager.run()
}

func (c *clientManager) run() {
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
				conn.writer <- message
				return true
			})
		}
	}
}

type clientManager struct {
	sync.RWMutex
	Clients    sync.Map
	HostName   string
	upgrader   websocket.HertzUpgrader
	handler    func(*websocket.Conn)
	broadcast  chan *types.Message
	register   chan *Client
	unregister chan *Client
}

func (c *clientManager) add(conn *Client) {
	c.Lock()
	defer c.Unlock()
	c.Clients.Store(conn.linkID, conn)
}

func (c *clientManager) del(conn *Client) {
	c.Lock()
	defer c.Unlock()
	c.Clients.Delete(conn.linkID)
}

func (c *clientManager) GetConn(linkID string) (bool, *Client) {
	conn, ok := c.Clients.Load(linkID)
	if ok {
		return true, conn.(*Client)
	}
	return false, nil
}
