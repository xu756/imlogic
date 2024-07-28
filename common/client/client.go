package client

import (
	"context"
	"imlogic/common/types"
	"imlogic/kitex_gen/im"
	"log"
	"os"
	"sync"
	"time"

	"github.com/hertz-contrib/websocket"
)

type Client struct {
	sync.RWMutex
	ctx        context.Context
	cancelFunc context.CancelFunc
	LinkId     string // websocket 连接 id
	UserId     int64  // 用户id
	hostName   string
	ws         *websocket.Conn
	send       chan *types.Message
	heartbeat  *time.Ticker
	MetaMsg    func(*Client, *im.MetaMsg)
	Logic      func(*Client, *types.Message)
	OnConnect  func(*Client)
	OnClose    func(*Client)
}

func NewClient(ctx context.Context, ws *websocket.Conn, linkID string, userId int64,
	heartbeat time.Duration,
	OnConnect func(*Client),
	OnClose func(*Client),
	MetaMsg func(*Client, *im.MetaMsg),
	Msglogic func(*Client, *types.Message),

) *Client {
	ctx, cancelFunc := context.WithCancel(ctx)
	ws.SetReadLimit(1024 * 1024 * 100)
	hostname, _ := os.Hostname()
	conn := &Client{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		hostName:   hostname,
		LinkId:     linkID,
		UserId:     userId,
		send:       make(chan *types.Message, 1024),
		heartbeat:  time.NewTicker(60 * time.Second),
		OnConnect:  OnConnect,
		OnClose:    OnClose,
		Logic:      Msglogic,
		MetaMsg:    MetaMsg,
	}
	go conn.setWs(ws)
	conn.OnConnect(conn)
	return conn
}

func (c *Client) Listen() {
	// go c.listenAndWrite()
	c.readMessage()

}

// listenAndRead 监听并读取消息
func (c *Client) readMessage() {
	for {
		select {
		case <-c.ctx.Done():

			return
		default:
			ws := c.getWs()
			if ws == nil {
				return
			}
			msg := &types.Message{}
			err := c.ws.ReadJSON(msg)
			if err != nil {
				return
			}
			go c.Logic(c, msg)
		}
	}

}

// listenAndWrite 监听并写入消息
// func (c *Client) listenAndWrite() {
// 	defer c.close()
// 	for {
// 		select {
// 		case <-c.ctx.Done():
// 			return
// 		case msg, ok := <-c.send:
// 			if !ok {
// 				return
// 			}
// 			go c.write(msg)
// 		case <-c.heartbeat.C:
// 			// go c.MetaMsg(c.heartbeatMsg())
// 		}

// 	}
// }

func (c *Client) SendMsg(msg *types.Message) {
	go c.write(msg)
}

// close 关闭连接
func (c *Client) close() {
	c.Lock()
	ws := c.getWs()
	if ws != nil {
		c.OnClose(c)
		c.ws.WriteMessage(websocket.CloseMessage, []byte{})
		c.cancelFunc()
		if err := c.ws.Close(); err != nil {
			log.Print("close:", err)
		}
		c.setWs(nil)

	}
	c.Unlock()
}

func (c *Client) getWs() *websocket.Conn {
	c.Lock()
	defer c.Unlock()
	return c.ws
}

func (c *Client) setWs(ws *websocket.Conn) {
	c.Lock()
	c.ws = ws
	c.Unlock()
	if ws == nil {
		return
	}
}

func (c *Client) write(msg *types.Message) {
	c.Lock()
	newMsg := &types.Message{}
	newMsg = msg
	// newMsg.LinkId = c.LinkId

	err := c.ws.WriteJSON(newMsg)
	if err != nil {
		c.close()
		c.Unlock()
		return
	}
	c.Unlock()
}

// connectMsg 连接消息
func (c *Client) ConnectMsg() *im.MetaMsg {
	return &im.MetaMsg{
		LinkId:   c.LinkId,
		UserId:   c.UserId,
		Status:   im.WsStatus_Connect,
		HostName: c.hostName,
	}
}

// heartbeatMsg 心跳消息
func (c *Client) HeartbeatMsg() *im.MetaMsg {
	return &im.MetaMsg{
		LinkId:   c.LinkId,
		Status:   im.WsStatus_Heartbeat,
		HostName: c.hostName,
	}
}

// disconnectMsg 断开连接消息
func (c *Client) DisconnectMsg() *im.MetaMsg {
	return &im.MetaMsg{
		LinkId:   c.LinkId,
		Status:   im.WsStatus_Disconnect,
		HostName: c.hostName,
	}
}
