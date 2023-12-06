package handler

import (
	"github.com/xu756/imlogic/internal/tool"
)

func (c *Client) logic(msg *Message) {
	ClientManager.broadcast <- &Message{
		MsgType: 100,
		Msg:     tool.TimeNowString(),
	}
}
