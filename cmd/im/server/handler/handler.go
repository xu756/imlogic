package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
	"github.com/xu756/imlogic/cmd/im/server/rpc"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/result"
	"github.com/xu756/imlogic/internal/tool"
	"github.com/xu756/imlogic/kitex_gen/im"
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
	err := ClientManager.upgrader.Upgrade(c, func(ws *websocket.Conn) {
		// todo 获取用户信息
		client := NewClient(ctx, ws, "admin", uuid.NewString(), "pc")
		msg, err := rpc.ImSrvClient.MetaMsg(ctx, &im.Message{
			Timestamp: tool.TimeNowUnixMilli(),
			Action:    "send",
			UserId:    client.userId,
			Hostname:  ClientManager.HostName,
			Device:    client.device,
			From:      client.linkID,
			To:        "im-rpc",
			MsgType:   "meta",
			MsgMeta:   &im.MsgMeta{DetailType: "connect", Version: config.GetVersion()},
		})
		if err != nil || msg.Success == false {
			result.HttpError(c, err)
			return
		}
		ClientManager.register <- client
		go client.listenAndWrite()
		client.listenAndRead()

	})
	if err != nil {
		result.HttpError(c, err)
		return
	}

}
