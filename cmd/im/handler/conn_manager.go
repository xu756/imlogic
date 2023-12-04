package handler

import (
	"github.com/hertz-contrib/websocket"
	"log"
	"sync"
)

var ConnManager = new(connManager)

func NewConnManager(HostName string) {
	ConnManager = &connManager{
		HostName: HostName,
		upgrader: websocket.HertzUpgrader{},
		handler:  ImHandler,
	}
}

func ImHandler(conn *websocket.Conn) {
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

type connManager struct {
	List     sync.Map
	HostName string
	upgrader websocket.HertzUpgrader
	handler  func(*websocket.Conn)
}

func (c *connManager) AddConn(conn *Conn) {
	c.List.Store(conn.linkID, conn)
}

func (c *connManager) DelConn(conn *Conn) {
	c.List.Delete(conn.linkID)
}

func (c *connManager) GetConn(linkID string) *Conn {
	conn, ok := c.List.Load(linkID)
	if ok {
		return conn.(*Conn)
	}
	return nil
}
