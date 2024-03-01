package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
	"github.com/xu756/imlogic/cmd/im/server/rpc"
	"github.com/xu756/imlogic/internal/tool"
	"github.com/xu756/imlogic/kitex_gen/im"
	"log"
	"sync"
	"time"
)

var (
	pongWait   = 60 * time.Second // 测试 暂时设置为 4s
	pingPeriod = (pongWait * 9) / 10
)

//// Message  消息结构体
//// todo 测试消息结构体，需要重新改
//type Message struct {
//	MsgType uint   `json:"msgType"`
//	Msg     string `json:"msg"`
//}

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
	client, err := rpc.ImSrvClient.Receive(c.ctx)
	if err != nil {
		client.Close()
		return
	}
	defer func() {
		c.close()
		client.Close()
	}()
	c.isOpen = true
	c.ws.SetReadLimit(1024 * 1024 * 100)
	//  设置读取超时时间 心跳
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
			rpcMsg := &im.Message{
				MsgId:     msg.MsgId,
				Device:    msg.Device,
				Timestamp: msg.Timestamp,
				Params:    msg.Params,
				Action:    msg.Action,
				From:      msg.From,
				To:        msg.To,
				MsgType:   msg.MsgType,
				MsgMeta: &im.MsgMeta{
					DetailType: msg.MsgMeta.DetailType,
					Version:    msg.MsgMeta.Version,
					Interval:   msg.MsgMeta.Interval,
				},
				MsgContent: &im.MsgContent{
					DetailType: msg.MsgContent.DetailType,
					Text:       msg.MsgContent.Text,
					ImgUrl:     msg.MsgContent.ImgUrl,
					AudioUrl:   msg.MsgContent.AudioUrl,
					VideoUrl:   msg.MsgContent.VideoUrl,
				},
			}
			err = client.Send(rpcMsg)
			if err != nil {
				client.Close()
				c.close()
			}
			recMsg, err := client.Recv()
			if err != nil {
				log.Print(err)
				return
			}
			go c.logic(RpcMsgToMsg(recMsg))
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
			//err := c.ws.WriteJSON("ping")
			//if err != nil {
			//	c.close()
			//	return
			//}
			c.writer <- &Message{
				MsgId:     uuid.NewString(),
				Device:    "pc", //todo 判断设备类型
				Timestamp: tool.TimeNowUnixMilli(),
				MsgType:   "meta",
				MsgMeta: MsgMeta{
					DetailType: "heartbeat",
					Version:    "1.0",
					Interval:   30000,
				},
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
		c.cancel() // 取消上下文  listenAndRead listenAndWrite 同时退出
		if err := c.ws.Close(); err != nil {
			log.Print(err)
		}
		c.isOpen = false
	}
}
