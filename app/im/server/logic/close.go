package logic

import (
	"imlogic/common/client"
	"imlogic/common/types"
	"imlogic/internal/tool"
	"log"

	"github.com/google/uuid"
)

func onClose(c *client.Client) {
	service.hub.DelOneClient(c)
	log.Print("onclose")
}

// 连接时
func onConnect(conn *client.Client) {
	log.Print("onconnect")
	conn.SendMsg(&types.Message{
		LinkId:    conn.LinkId,
		Sender:    conn.UserId,
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
