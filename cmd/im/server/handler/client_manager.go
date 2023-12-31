package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"os"
	"sync"
)

var ClientManager = new(clientManager)

func NewClientnManager() {
	hostname, _ := os.Hostname()

	ClientManager = &clientManager{
		HostName:   hostname,
		broadcast:  make(chan *Message),
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
				select {
				case conn.writer <- message:
				default:
					c.del(conn)
				}
				return true
			})
		}
	}
}

type clientManager struct {
	Clients     sync.Map
	HostName    string
	upgrader    websocket.HertzUpgrader
	handler     func(*websocket.Conn)
	clientsLock sync.RWMutex
	broadcast   chan *Message
	register    chan *Client
	unregister  chan *Client
}

func (c *clientManager) add(conn *Client) {
	c.Clients.Store(conn.linkID, conn)
}

func (c *clientManager) del(conn *Client) {
	c.clientsLock.Lock()
	defer c.clientsLock.Unlock()
	c.Clients.Delete(conn.linkID)
}

func (c *clientManager) GetConn(linkID string) (bool, *Client) {
	conn, ok := c.Clients.Load(linkID)
	if ok {
		return true, conn.(*Client)
	}
	return false, nil
}
