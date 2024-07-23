package client

import (
	"context"
	"imlogic/common/config"
	"imlogic/common/types"
	"imlogic/kitex_gen/im"
	"log"
	"sync"
	"time"

	"github.com/hertz-contrib/websocket"
)

type Client struct {
	sync.RWMutex
	Ctx       context.Context
	Cancel    context.CancelFunc
	LinkId    string // websocket 连接 id
	UserId    int64  // 用户id
	Ws        *websocket.Conn
	IsOpen    bool
	Send      chan *types.Message
	Heartbeat *time.Ticker
	MetaMsg   func(msg *im.MetaMsg)
	Logic     func(*Client, *types.Message)
	OnClose   func(*Client)
}

// listenAndRead 监听并读取消息
func (c *Client) ListenAndRead() {
	defer func() {
		c.Close()
	}()
	for {
		select {
		case <-c.Ctx.Done():
			return
		default:
			msg := &types.Message{}
			err := c.Ws.ReadJSON(msg)
			if err != nil {
				return
			}
			go c.Logic(c, msg)
		}
	}

}

// listenAndWrite 监听并写入消息
func (c *Client) ListenAndWrite() {
	defer c.Close()
	for {
		select {
		case <-c.Ctx.Done():
			return
		case msg, ok := <-c.Send:
			if !ok {
				return
			}
			c.Write(msg)
		case <-c.Heartbeat.C:
			// go c.MetaMsg(c.heartbeatMsg())
		}

	}
}

// close 关闭连接
func (c *Client) Close() {
	c.Lock()
	defer c.Unlock()
	if c.IsOpen {
		c.OnClose(c)
		c.Ws.WriteMessage(websocket.CloseMessage, []byte{})
		go c.MetaMsg(c.DisconnectMsg())
		c.Cancel()
		if err := c.Ws.Close(); err != nil {
			log.Print("close:", err)
		}
		c.IsOpen = false
	}
}

func (c *Client) Write(msg *types.Message) {
	msg.LinkId = c.LinkId
	msg.MsgMeta.Version = config.GetVersion()
	err := c.Ws.WriteJSON(msg)
	if err != nil {
		c.Close()
		return
	}
}

// connectMsg 连接消息
func (c *Client) ConnectMsg() *im.MetaMsg {
	return &im.MetaMsg{
		LinkId: c.LinkId,
		UserId: c.UserId,
		Status: im.WsStatus_Connect,
	}
}

// heartbeatMsg 心跳消息
func (c *Client) HeartbeatMsg() *im.MetaMsg {
	return &im.MetaMsg{
		LinkId: c.LinkId,
		Status: im.WsStatus_Heartbeat,
	}
}

// disconnectMsg 断开连接消息
func (c *Client) DisconnectMsg() *im.MetaMsg {
	return &im.MetaMsg{
		LinkId: c.LinkId,
		Status: im.WsStatus_Disconnect,
	}
}
