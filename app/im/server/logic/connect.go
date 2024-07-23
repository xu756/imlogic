package logic

import (
	"context"
	"imlogic/common/client"
	"imlogic/common/types"
	"imlogic/internal/result"
	"imlogic/internal/tool"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
)

var (
	pongWait   = 60 * time.Second // 测试 暂时设置为 4s
	pingPeriod = (pongWait * 9) / 10
)

// NewClient 创建一个新的连接
func NewClient(ctx context.Context, ws *websocket.Conn, linkID, device string, userId int64) *client.Client {
	ctx, cancel := context.WithCancel(ctx)
	return &client.Client{
		Ctx:       ctx,
		Cancel:    cancel,
		Ws:        ws,
		IsOpen:    true,
		LinkId:    linkID,
		UserId:    userId,
		Device:    device,
		Send:      make(chan *types.Message, 1024),
		Heartbeat: time.NewTicker(pingPeriod),
		OnClose:   Close,
		Logic:     Msglogic,
	}
}

// 发送连接消息
func initClient(c *client.Client) {
	c.IsOpen = true
	c.Ws.SetReadLimit(1024 * 1024 * 100)
	err := c.Ws.WriteJSON(types.Message{
		LinkId:    c.LinkId,
		MsgId:     uuid.NewString(),
		Timestamp: tool.TimeNowUnixMilli(),
		ChatType:  types.SystemMessage,
		MsgMeta: types.MsgMeta{
			Detail:  "connect",
			Version: "1.0",
		},
	})
	if err != nil {
		c.Close()
		return
	}

}

func Connect(ctx context.Context, c *app.RequestContext) {
	err := service.hub.UpgradeOneWs(c, func(ws *websocket.Conn) {
		// todo 获取用户信息
		wsClient := NewClient(ctx, ws, uuid.NewString(), "pc", 0)
		MetaMsg(ctx, wsClient, wsClient.ConnectMsg())
		initClient(wsClient)
		service.hub.Register <- wsClient
		go wsClient.ListenAndWrite()
		wsClient.ListenAndRead()
	})

	if err != nil {
		result.HttpError(c, err)
		return
	}

}
