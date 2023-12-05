package handler

import "log"

func (c *Client) logic(msg *Message) {
	log.Print(msg)
}
