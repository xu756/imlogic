package handler

import (
	"context"
	"log"
	"sync"
	"time"

	"imlogic/app/im/server/rpc"
	"imlogic/common/config"
	"imlogic/common/types"
	"imlogic/internal/tool"
	"imlogic/kitex_gen/im"

	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
)

var (
	pongWait   = 60 * time.Second // 测试 暂时设置为 4s
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
	linkID    string // websocket 连接 id
	userId    string // 用户id
	device    string // 设备类型
	ws        *websocket.Conn
	isOpen    bool
	send      chan *types.Message
	heartbeat *time.Ticker
}

// NewClient 创建一个新的连接
func NewClient(ctx context.Context, ws *websocket.Conn, userId, linkID, device string) *Client {
	ctx, cancel := context.WithCancel(ctx)
	return &Client{
		ctx:       ctx,
		cancel:    cancel,
		ws:        ws,
		isOpen:    true,
		linkID:    linkID,
		userId:    userId,
		device:    device,
		send:      make(chan *types.Message, 1024),
		heartbeat: time.NewTicker(pingPeriod),
	}
}

// listenAndRead 监听并读取消息
func (c *Client) listenAndRead() {
	defer func() {
		c.close()
	}()
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
			rpcMsg := &im.Message{
				MsgId:     uuid.NewString(),
				LinkId:    c.linkID,
				Hostname:  hub.HostName,
				Timestamp: tool.TimeNowUnixMilli(),
				Device:    c.device,
				Action:    msg.Action,
				From:      c.userId,
				To:        msg.To,
				MsgType:   msg.MsgType,
				MsgMeta: &im.MsgMeta{
					DetailType: msg.MsgMeta.DetailType,
					Version:    msg.MsgMeta.Version,
				},
				MsgContent: &im.MsgContent{
					DetailType: msg.MsgContent.DetailType,
					Text:       msg.MsgContent.Text,
					ImgUrl:     msg.MsgContent.ImgUrl,
					AudioUrl:   msg.MsgContent.AudioUrl,
					VideoUrl:   msg.MsgContent.VideoUrl,
				},
				Params: msg.Params,
			}
			go func() {
				receive, err := rpc.ImSrvClient.Receive(c.ctx, rpcMsg)
				if err != nil {
					return
				}
				c.send <- types.RpcMsgToMsg(receive)
			}()
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
			c.Write(msg)
		case <-c.heartbeat.C:
			go c.rpcMetaMsg(c.heartbeatMsg())
		}

	}
}

// close 关闭连接
func (c *Client) close() {
	c.Lock()
	defer c.Unlock()
	if c.isOpen {
		hub.unregister <- c
		c.ws.WriteMessage(websocket.CloseMessage, []byte{})
		_, _ = rpc.ImSrvClient.MetaMsg(c.ctx, c.disconnectMsg())
		c.cancel() // 取消上下文  listenAndRead listenAndWrite 同时退出
		if err := c.ws.Close(); err != nil {
			log.Print("close:", err)
		}
		c.isOpen = false
	}
}

func (c *Client) Write(msg *types.Message) {
	msg.LinkId = c.linkID
	msg.MsgMeta.Version = config.GetVersion()
	err := c.ws.WriteJSON(msg)
	if err != nil {
		c.close()
		return
	}
}

// connectMsg 连接消息
func (c *Client) connectMsg() *im.Message {
	return &im.Message{
		Timestamp: tool.TimeNowUnixMilli(),
		MsgId:     uuid.NewString(),
		Action:    "send",
		LinkId:    c.linkID,
		Hostname:  hub.HostName,
		Device:    c.device,
		From:      c.userId,
		To:        "im-rpc",
		MsgType:   "meta",
		MsgMeta:   &im.MsgMeta{DetailType: "connect", Version: config.GetVersion()},
	}
}

// heartbeatMsg 心跳消息
func (c *Client) heartbeatMsg() *im.Message {
	return &im.Message{
		Timestamp: tool.TimeNowUnixMilli(),
		MsgId:     uuid.NewString(),
		Device:    c.device,
		Hostname:  hub.HostName,
		From:      c.userId,
		To:        "im-rpc",
		LinkId:    c.linkID,
		MsgType:   "meta",
		MsgMeta: &im.MsgMeta{
			DetailType: "heartbeat",
			Version:    config.GetVersion(),
		},
	}
}

// disconnectMsg 断开连接消息
func (c *Client) disconnectMsg() *im.Message {
	return &im.Message{
		Device:    c.device,
		MsgId:     uuid.NewString(),
		Timestamp: tool.TimeNowUnixMilli(),
		From:      c.userId,
		LinkId:    c.linkID,
		Hostname:  hub.HostName,
		To:        "im-rpc",
		MsgType:   "meta",
		MsgMeta: &im.MsgMeta{
			DetailType: "disconnect",
			Version:    config.GetVersion(),
		},
	}
}

// rpcMetaMsg
func (c *Client) rpcMetaMsg(msg *im.Message) {
	res, err := rpc.ImSrvClient.MetaMsg(c.ctx, msg)
	if err != nil || res.Success == false {
		return
	}
	c.send <- &types.Message{
		MsgId: res.MsgId,
		From:  res.From,
		To:    res.To,
	}
}
