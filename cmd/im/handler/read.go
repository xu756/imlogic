package handler

import "github.com/xu756/imlogic/internal/tool"

func (c *Client) logic(msg *Message) {
	if ok, client := ClientManager.GetConn("test"); ok {
		client.writer <- &Message{
			MsgType: 100,
			Msg:     tool.TimeNowString(),
		}
	}
}
