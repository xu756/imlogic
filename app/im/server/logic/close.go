package logic

import (
	"imlogic/common/client"
	"imlogic/internal/tool"
	"imlogic/kitex_gen/base"

	"github.com/google/uuid"
)

func onClose(c *client.Client) {
	service.hub.DelOneClient(c)
	MetaMsg(c, c.DisconnectMsg())

}

// 连接时
func onConnect(conn *client.Client) {
	conn.SendMsg(&base.Message{
		LinkId:    conn.GetLinkId(),
		Sender:    conn.GetUserId(),
		MsgId:     uuid.NewString(),
		Timestamp: tool.TimeNowUnixMilli(),
		ChatType:  base.ChatType_SystemMessage,
		Content:   "连接成功",
	})
	MetaMsg(conn, conn.ConnectMsg())

}
