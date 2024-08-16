package logic

import (
	"context"
	"imlogic/common/client"
	"imlogic/internal/result"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
)

var (
	pongWait = 60 * time.Second // 测试 暂时设置为 4s
)

func Connect(ctx context.Context, c *app.RequestContext) {
	userInfo, err := service.Jwt.GetUserInfoFromCookieToken(c)
	if err != nil {
		result.HttpError(c, err)
		return
	}
	err = service.hub.UpgradeOneWs(c, func(ws *websocket.Conn) {
		conn := client.NewClient(
			ctx, ws,
			uuid.NewString(),
			userInfo.UserId,
			pongWait,
			onConnect,
			onClose,
			MetaMsg,
			MsgLogic,
		)
		service.hub.AddOneClient(conn)
		conn.Listen()
	})
	if err != nil {
		log.Print(err)
		result.HttpError(c, err)
		return
	}

}
