package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
	"github.com/xu756/imlogic/cmd/im/server/rpc"
	"github.com/xu756/imlogic/common/types"
	"github.com/xu756/imlogic/internal/tool"
	"github.com/xu756/imlogic/kitex_gen/im"
	"log"
	"time"
)

var (
	pongWait   = 5 * time.Second // 测试 暂时设置为 4s
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
	linkID string // websocket 连接 id
	userId string // 用户id
	device string // 设备类型
	ws     *websocket.Conn
	isOpen bool
	send   chan *types.Message
}

// NewClient 创建一个新的连接
func NewClient(ctx context.Context, ws *websocket.Conn, userId, linkID, device string) *Client {
	ctx, cancel := context.WithCancel(ctx)
	return &Client{
		ctx:    ctx,
		cancel: cancel,
		ws:     ws,
		linkID: linkID,
		userId: userId,
		device: device,
		send:   make(chan *types.Message, 1024),
	}
}

// listenAndRead 监听并读取消息
func (c *Client) listenAndRead() {
	defer c.close()
	//  设置读取超时时间 心跳

	for {
		msg := &types.Message{}
		err := c.ws.ReadJSON(msg)
		if err != nil {
			return
		}
		msg.From = c.linkID
		c.logic(msg)

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
			c.Write(msg)
		}

	}
}

// close 关闭连接
func (c *Client) close() {

	if c.isOpen {
		hub.unregister <- c
		close(c.send)
		c.ws.WriteMessage(websocket.CloseMessage, []byte{})
		_, _ = rpc.ImSrvClient.MetaMsg(c.ctx, &im.Message{
			UserId:    c.userId,
			Hostname:  hub.HostName,
			Device:    c.device,
			Timestamp: tool.TimeNowUnixMilli(),
			Action:    "send",
			From:      c.linkID,
			To:        "im-rpc",
			MsgType:   "meta",
			MsgMeta:   &im.MsgMeta{DetailType: "disconnect", Version: "1.0"},
		})
		hub.del(c)
		c.cancel() // 取消上下文  listenAndRead listenAndWrite 同时退出
		if err := c.ws.Close(); err != nil {
			log.Print(err)
		}
		c.isOpen = false
	}
}

func (c *Client) Write(msg *types.Message) {
	msg.MsgId = uuid.NewString()
	msg.To = c.linkID
	err := c.ws.WriteJSON(msg)
	if err != nil {
		c.close()
		return
	}
}
