package handler

import (
	"imlogic/common/types"
)

func (c *Client) logic(msg *types.Message) {
	hub.broadcast <- msg
}
