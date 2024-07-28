package logic

import (
	"imlogic/common/client"
	"imlogic/common/types"
	"imlogic/internal/tool"

	"github.com/google/uuid"
)

func onClose(c *client.Client) {
	service.hub.DelOneClient(c)
	MetaMsg(c, c.DisconnectMsg())

}

// 连接时
func onConnect(conn *client.Client) {
	conn.SendMsg(&types.Message{
		LinkId:    conn.LinkId,
		Sender:    conn.UserId,
		MsgId:     uuid.NewString(),
		Timestamp: tool.TimeNowUnixMilli(),
		ChatType:  types.SystemMessage,
		Content:   "连接成功",
	})
	MetaMsg(conn, conn.ConnectMsg())

}
