package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
	"github.com/xu756/imlogic/cmd/im/server/rpc"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/common/types"
	"github.com/xu756/imlogic/internal/tool"
	"github.com/xu756/imlogic/kitex_gen/im"
	"github.com/xu756/imlogic/kitex_gen/im/imsrv"
	"log"
	"sync"
	"time"
)

var (
	pongWait   = 50 * time.Second // 测试 暂时设置为 4s
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
	rpcClient imsrv.ImSrv_ReceiveClient
	linkID    string // websocket 连接 id
	userId    string // 用户id
	device    string // 设备类型
	ws        *websocket.Conn
	isOpen    bool
	send      chan *types.Message
	rpcSend   chan *types.Message
}

// NewClient 创建一个新的连接
func NewClient(ctx context.Context, ws *websocket.Conn, userId, linkID, device string) *Client {
	ctx, cancel := context.WithCancel(ctx)
	rpcClient, err := rpc.ImSrvClient.Receive(ctx)
	if err != nil {
		panic(err)
	}
	return &Client{
		ctx:       ctx,
		cancel:    cancel,
		ws:        ws,
		isOpen:    true,
		rpcClient: rpcClient,
		linkID:    linkID,
		userId:    userId,
		device:    device,
		send:      make(chan *types.Message, 1024),
	}
}

func (c *Client) listenRpcMsg() {
	defer func() {
		c.close()
		c.rpcClient.Close()
	}()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			recMsg, err := c.rpcClient.Recv()
			if err != nil {
				return
			}
			hub.broadcast <- types.RpcMsgToMsg(recMsg)
		}
	}
}

// listenAndRead 监听并读取消息
func (c *Client) listenAndRead() {

	defer func() {
		c.close()
		c.rpcClient.Close()
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
				UserId:    c.userId,
				Hostname:  hub.HostName,
				Timestamp: msg.Timestamp,
				Device:    msg.Device,
				Action:    msg.Action,
				From:      c.linkID,
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
			err = c.rpcClient.Send(rpcMsg)
			if err != nil {
				return
			}

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
		c.rpcClient.Close()
		hub.del(c)
		c.cancel() // 取消上下文  listenAndRead listenAndWrite 同时退出
		if err := c.ws.Close(); err != nil {
			log.Print("close:", err)
		}
		c.isOpen = false
	}
}

func (c *Client) Write(msg *types.Message) {
	msg.To = c.linkID
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
		UserId:    c.userId,
		Hostname:  hub.HostName,
		Device:    c.device,
		From:      c.linkID,
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
		From:      c.linkID,
		To:        "im-rpc",
		UserId:    c.userId,
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
		From:      c.linkID,
		UserId:    c.userId,
		Hostname:  hub.HostName,
		To:        "im-rpc",
		MsgType:   "meta",
		MsgMeta: &im.MsgMeta{
			DetailType: "disconnect",
			Version:    config.GetVersion(),
		},
	}
}
