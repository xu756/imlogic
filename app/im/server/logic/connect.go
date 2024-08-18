package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/websocket"
	"imlogic/common/client"
	"imlogic/internal/result"
	"log"
)

var ()

func Connect(ctx context.Context, c *app.RequestContext) {
	userInfo, err := service.Jwt.GetUserInfoFromHeardToken(c)
	if err != nil {
		result.HttpError(c, err)
		return
	}
	err = service.hub.UpgradeOneWs(c, func(ws *websocket.Conn) {
		conn := client.NewClient(
			ctx, ws,
			uuid.NewString(),
			userInfo.UserId,
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
