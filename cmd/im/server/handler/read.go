package handler

func (c *Client) logic(msg *Message) {
	ClientManager.broadcast <- msg
}
