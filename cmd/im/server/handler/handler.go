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
	"log"
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
		// todo 获取
		linkId := uuid.NewString()
		userId := "admin"
		device := "pc"
		client := NewClient(ctx, ws, linkId)
		msg, err := rpc.ImSrvClient.MetaMsg(ctx, &im.Message{
			Device:    device,
			Timestamp: tool.TimeNowUnixMilli(),
			Action:    "send",
			Params: map[string]string{
				"userId": userId,
			},
			From:    linkId,
			To:      "im-rpc",
			MsgType: "meta",
			MsgMeta: &im.MsgMeta{DetailType: "connect", Version: "1.0", Interval: 0},
		})
		if err != nil || msg.Success == false {
			result.HttpError(c, err)
			return
		}
		go client.listenAndWrite()
		ClientManager.register <- client
		client.listenAndRead()
	})
	if err != nil {
		log.Print(err)
		result.HttpError(c, err)
		return
	}

}
