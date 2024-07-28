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
	ctx       context.Context
	cancel    context.CancelFunc
	LinkId    string // websocket 连接 id
	UserId    int64  // 用户id
	hostName  string
	ws        *websocket.Conn
	isOpen    bool
	send      chan *types.Message
	heartbeat *time.Ticker
	MetaMsg   func(*Client, *im.MetaMsg)
	Logic     func(*Client, *types.Message)
	OnConnect func(*Client)
	OnClose   func(*Client)
}

func NewClient(ctx context.Context, ws *websocket.Conn, linkID string, userId int64,
	heartbeat time.Duration,
	OnConnect func(*Client),
	OnClose func(*Client),
	MetaMsg func(*Client, *im.MetaMsg),
	Msglogic func(*Client, *types.Message),

) *Client {
	ctx, cancel := context.WithCancel(ctx)
	ws.SetReadLimit(1024 * 1024 * 100)
	hostname, _ := os.Hostname()
	conn := &Client{
		ctx:       ctx,
		cancel:    cancel,
		hostName:  hostname,
		ws:        ws,
		LinkId:    linkID,
		isOpen:    true,
		UserId:    userId,
		send:      make(chan *types.Message, 1024),
		heartbeat: time.NewTicker(60 * time.Second),
		OnConnect: OnConnect,
		OnClose:   OnClose,
		Logic:     Msglogic,
		MetaMsg:   MetaMsg,
	}
	conn.OnConnect(conn)
	return conn
}

func (c *Client) Listen() {
	// go c.listenAndWrite()
	c.listenAndRead()

}

// listenAndRead 监听并读取消息
func (c *Client) listenAndRead() {
	defer c.close()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
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
	defer c.Unlock()
	if c.isOpen {
		c.OnClose(c)
		c.ws.WriteMessage(websocket.CloseMessage, []byte{})
		go c.MetaMsg(c, c.DisconnectMsg())
		c.cancel()
		if err := c.ws.Close(); err != nil {
			log.Print("close:", err)
		}
		c.isOpen = false
	}
}

func (c *Client) write(msg *types.Message) {
	c.Lock()
	defer c.Unlock()
	msg.LinkId = c.LinkId
	err := c.ws.WriteJSON(msg)
	if err != nil {
		c.close()
		return
	}
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
