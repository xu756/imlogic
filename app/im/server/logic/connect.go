package logic

import (
	"context"
	"imlogic/common/client"
	"imlogic/common/types"
	"imlogic/internal/result"
	"imlogic/internal/tool"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
)

var (
	pongWait   = 60 * time.Second // 测试 暂时设置为 4s
	pingPeriod = (pongWait * 9) / 10
)

// 连接时
func onConnect(conn *client.Client) {
	log.Print("onconnect")
	conn.SendMsg(&types.Message{
		LinkId:    conn.LinkId,
		MsgId:     uuid.NewString(),
		Timestamp: tool.TimeNowUnixMilli(),
		ChatType:  types.SystemMessage,
		MsgMeta: types.MsgMeta{
			Detail:  "connect",
			Version: "1.0",
		},
	})
	MetaMsg(conn, conn.ConnectMsg())

}

func Logic(ws *websocket.Conn) {
	ctx := context.Background()
	conn := client.NewClient(ctx, ws,
		uuid.NewString(),
		0,
		60*time.Second,
		onConnect,
		onClose,
		MetaMsg, Msglogic)
	service.hub.AddOneClient(conn)
	go conn.ListenAndWrite()
	conn.ListenAndRead()

}

func Connect(ctx context.Context, c *app.RequestContext) {
	err := service.hub.UpgradeOneWs(c, Logic)
	if err != nil {
		log.Print(err)
		result.HttpError(c, err)
		return
	}

}
