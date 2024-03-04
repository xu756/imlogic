package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
	"github.com/xu756/imlogic/cmd/im/server/rpc"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/common/types"
	"github.com/xu756/imlogic/internal/result"
	"github.com/xu756/imlogic/internal/tool"
)

var HttpServer *server.Hertz

// InitRouter 路由
func InitRouter() {
	h := server.Default(
		server.WithHostPorts(config.RunData.Addr.ImAddr),
		server.WithReadBufferSize(1024*1024*100),
		server.WithMaxRequestBodySize(1024*1024*100),
	)
	//h.Use(middleware.HertzJwt())
	h.GET("/connect", connect)
	HttpServer = h
}

func connect(ctx context.Context, c *app.RequestContext) {
	err := hub.upgrader.Upgrade(c, func(ws *websocket.Conn) {
		// todo 获取用户信息
		client := NewClient(ctx, ws, "admin", uuid.NewString(), "pc")
		msg, err := rpc.ImSrvClient.MetaMsg(ctx, client.connectMsg())
		if err != nil || msg.Success == false {
			result.HttpError(c, err)
			return
		}
		initClient(client)
		hub.register <- client
		go client.listenAndWrite()
		go client.listenRpcMsg()
		client.listenAndRead()

	})
	if err != nil {
		result.HttpError(c, err)
		return
	}

}

func initClient(c *Client) {
	c.isOpen = true
	c.ws.SetReadLimit(1024 * 1024 * 100)
	err := c.ws.WriteJSON(types.Message{
		MsgId:     uuid.NewString(),
		Device:    c.device,
		Timestamp: tool.TimeNowUnixMilli(),
		From:      "im-rpc",
		To:        c.linkID,
		MsgType:   "meta",
		MsgMeta: types.MsgMeta{
			DetailType: "heartbeat",
			Version:    config.GetVersion(),
		},
	})
	if err != nil {
		c.close()
		return
	}

}
