package client

import (
	"context"
	"imlogic/kitex_gen/base"
	"os"
	"sync"
	"time"

	"github.com/hertz-contrib/websocket"
)

const (
	// 心跳时间
	HeartbeatTime = 30000
)

type Client struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	linkId     string // websocket 连接 id
	userId     int64  // 用户id
	hostName   string
	lock       sync.Mutex
	ws         *websocket.Conn
	send       chan *base.Message
	heartbeat  *time.Ticker
	metaMsg    func(*Client, *base.MetaMsg)
	logic      func(*Client, *base.Message)
	connect    func(*Client)
	OnClose    func(*Client)
}

func NewClient(ctx context.Context, ws *websocket.Conn, linkID string, userId int64,
	connect func(*Client),
	close func(*Client),
	metaMsg func(*Client, *base.MetaMsg),
	msgLogic func(*Client, *base.Message),

) *Client {
	ctx, cancelFunc := context.WithCancel(ctx)
	ws.SetReadLimit(1024 * 1024 * 100)
	hostname, _ := os.Hostname()
	conn := &Client{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		hostName:   hostname,
		linkId:     linkID,
		userId:     userId,
		send:       make(chan *base.Message, 1024),
		heartbeat:  time.NewTicker(HeartbeatTime * time.Second),
		connect:    connect,
		OnClose:    close,
		logic:      msgLogic,
		metaMsg:    metaMsg,
	}
	go conn.setWs(ws)
	conn.connect(conn)
	return conn
}

func (c *Client) Listen() {
	go c.listenAndWrite()
	c.readMessage()
}

// listenAndRead 监听并读取消息
func (c *Client) readMessage() {
	defer c.close()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			if c.ws == nil {
				c.cancelFunc()
				return
			}
			msg := &base.Message{}
			err := c.ws.ReadJSON(msg)
			if err != nil {
				return
			}
			msg.Sender = c.GetUserId()
			go c.logic(c, msg)
		}
	}

}

// listenAndWrite 监听并写入消息
func (c *Client) listenAndWrite() {
	defer c.close()
	for {
		select {
		case <-c.ctx.Done():
			return
		case msg, ok := <-c.send:
			if !ok {
				return
			}
			go c.write(msg)
		case <-c.heartbeat.C:
			c.close()
		}
	}
}

// 重新设置heartbeat
func (c *Client) ResetHeartbeat() {
	c.heartbeat.Reset(HeartbeatTime * time.Second)
}

func (c *Client) SendMsg(msg *base.Message) {
	c.send <- msg
}

// close 关闭连接
func (c *Client) close() {
	c.lock.Lock()
	if c.ws != nil {
		c.OnClose(c)
		c.ws.WriteMessage(websocket.CloseMessage, []byte{})
		c.cancelFunc()
		c.ws.Close()
		c.setWs(nil)
	}
	c.lock.Unlock()
}

func (c *Client) setWs(ws *websocket.Conn) {
	c.lock.Lock()
	c.ws = ws
	c.lock.Unlock()
	if ws == nil {
		return
	}
}

func (c *Client) write(msg *base.Message) {
	c.lock.Lock()
	if c.ws == nil {
		c.lock.Unlock()
		return
	}
	err := c.ws.WriteJSON(msg)
	if err != nil {
		c.close()
		c.lock.Unlock()
		return
	}
	c.lock.Unlock()
}

// connectMsg 连接消息
func (c *Client) ConnectMsg() *base.MetaMsg {
	return &base.MetaMsg{
		LinkId:   c.GetLinkId(),
		UserId:   c.GetUserId(),
		Status:   base.WsStatus_Connect,
		HostName: c.hostName,
	}
}

// heartbeatMsg 心跳消息
func (c *Client) HeartbeatMsg() *base.MetaMsg {
	return &base.MetaMsg{
		LinkId:   c.GetLinkId(),
		UserId:   c.GetUserId(),
		Status:   base.WsStatus_Heartbeat,
		HostName: c.hostName,
	}
}

// disconnectMsg 断开连接消息
func (c *Client) DisconnectMsg() *base.MetaMsg {
	return &base.MetaMsg{
		LinkId:   c.GetLinkId(),
		UserId:   c.GetUserId(),
		Status:   base.WsStatus_Disconnect,
		HostName: c.hostName,
	}
}

// GetLinkId 获取连接id
func (c *Client) GetLinkId() string {
	return c.linkId
}

// GetUserId 获取用户id
func (c *Client) GetUserId() int64 {
	return c.userId
}
