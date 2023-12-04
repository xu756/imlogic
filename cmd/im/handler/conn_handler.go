package handler

import (
	"context"
	"github.com/hertz-contrib/websocket"
	"sync"
)

type Message struct {
	MsgType uint8  `json:"msg_type"`
	Msg     string `json:"msg"`
}

type Conn struct {
	ctx    context.Context
	mutex  sync.Mutex
	linkID string // websocket 连接 id
	ws     *websocket.Conn
	isOpen bool
	reader chan Message
	Writer chan Message
	Close  chan bool
}

// NewConn 创建一个新的连接
func NewConn(ctx context.Context, ws *websocket.Conn, linkID string) *Conn {
	return &Conn{
		ctx:    ctx,
		ws:     ws,
		linkID: linkID,
		reader: make(chan Message, 1024),
		Writer: make(chan Message, 1024),
		Close:  make(chan bool),
	}
}

// listenAndRead 监听并读取消息
func (c *Conn) listenAndRead() {
	// 通过ctx控制读取
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			var msg Message
			err := c.ws.ReadJSON(&msg)
			if err != nil {
				// todo 错误处理
				return
			}
			c.reader <- msg
		}
	}
}

// listenAndWrite 监听并写入消息
func (c *Conn) listenAndWrite() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case msg := <-c.Writer:
			err := c.ws.WriteJSON(msg)
			if err != nil {
				// todo 错误处理
				return
			}
		}
	}
}
