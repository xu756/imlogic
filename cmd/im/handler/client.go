package handler

import (
	"context"
	"github.com/hertz-contrib/websocket"
	"log"
	"sync"
	"time"
)

var (
	writeWait  = 10 * time.Second
	pongWait   = 4 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

type Message struct {
	MsgType uint   `json:"msgType"`
	Msg     string `json:"msg"`
}

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
	mutex  sync.Mutex
	linkID string // websocket 连接 id
	ws     *websocket.Conn
	isOpen bool
	writer chan *Message
}

// NewClient 创建一个新的连接
func NewClient(ctx context.Context, ws *websocket.Conn, linkID string) *Client {
	ctx, cancel := context.WithCancel(ctx)
	client := &Client{
		ctx:    ctx,
		cancel: cancel,
		ws:     ws,
		linkID: linkID,
		//reader: make(chan *Message, 1024),
		writer: make(chan *Message, 1024),
	}
	return client
}

// listenAndRead 监听并读取消息
func (c *Client) listenAndRead() {
	defer func() {
		c.close()
	}()
	c.isOpen = true
	c.ws.SetReadLimit(1024 * 1024 * 100)
	c.ws.SetReadDeadline(time.Now().Add(pongWait)) // 设置读取超时时间
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			var msg *Message
			err := c.ws.ReadJSON(&msg)
			if err != nil {
				c.close()
				return
			}
			go c.logic(msg)
		}
	}
}

// listenAndWrite 监听并写入消息
func (c *Client) listenAndWrite() {
	ticker := time.NewTicker(pingPeriod) // 定时发送心跳
	defer func() {
		ticker.Stop()
		c.close()
	}()
	c.ws.WriteJSON(map[string]interface{}{
		"msgType": 100,
		"msg":     "连接成功",
		"linkID":  c.linkID,
	})
	for {
		select {
		case <-c.ctx.Done():
			return
		case msg := <-c.writer:
			err := c.ws.WriteJSON(msg)
			if err != nil {
				c.close()
				return
			}
		case <-ticker.C:
			err := c.ws.WriteJSON("ping")
			if err != nil {
				c.close()
				return
			}
		}

	}
}

// close 关闭连接
func (c *Client) close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.isOpen {
		ClientManager.unregister <- c
		c.cancel() // 取消上下文
		if err := c.ws.Close(); err != nil {
			log.Print(err)
		}
		c.isOpen = false
	}
}
