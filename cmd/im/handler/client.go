package handler

import (
	"context"
	"github.com/hertz-contrib/websocket"
	"log"
	"sync"
)

type Message struct {
	MsgType uint8  `json:"msg_type"`
	Msg     string `json:"msg"`
}

type Client struct {
	ctx    context.Context
	mutex  sync.Mutex
	linkID string // websocket 连接 id
	ws     *websocket.Conn
	isOpen bool
	//reader chan *Message
	writer chan *Message
	Close  chan bool
}

// NewClient 创建一个新的连接
func NewClient(ctx context.Context, ws *websocket.Conn, linkID string) *Client {
	client := &Client{
		ctx:    ctx,
		ws:     ws,
		linkID: linkID,
		//reader: make(chan *Message, 1024),
		writer: make(chan *Message, 1024),
		Close:  make(chan bool),
	}
	return client
}

// listenAndRead 监听并读取消息
func (c *Client) listenAndRead() {
	// 通过ctx控制读取
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			msg := new(Message)
			mt, message, err := c.ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			msg.MsgType = uint8(mt)
			msg.Msg = string(message)
			log.Print(msg)
			//go c.logic(msg)
		}
	}
}

// listenAndWrite 监听并写入消息
func (c *Client) listenAndWrite() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case msg := <-c.writer:
			err := c.ws.WriteJSON(msg)
			if err != nil {
				// todo 错误处理
				return
			}
		}
	}
}
