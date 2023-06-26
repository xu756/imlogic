package logic

import (
	"context"
	"nhooyr.io/websocket"
	"sync"
)

type Client struct {
	connId     string
	lock       sync.Mutex
	ctx        context.Context
	cancelFunc context.CancelFunc
	ws         *websocket.Conn
}

func (c *Client) Close(code websocket.StatusCode, reason string) error {
	c.cancelFunc()
	// 关闭之前先判断ws是否存在
	ws := c.getWs()
	if ws != nil {
		err := ws.Close(code, reason)
		Hubs.unregister <- c
		c.setWs(nil)
		return err
	}
	return nil
}
func (c *Client) setWs(ws *websocket.Conn) error {
	c.lock.Lock()
	c.ws = ws
	c.lock.Unlock()
	if ws == nil {
		return nil
	}
	// 读消息
	go c.readMessage()
	return nil
}

func (c *Client) getWs() *websocket.Conn {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.ws
}
