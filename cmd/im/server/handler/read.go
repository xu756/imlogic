package handler

import "github.com/xu756/imlogic/common/types"

func (c *Client) logic(msg *types.Message) {
	ClientManager.broadcast <- msg
}
