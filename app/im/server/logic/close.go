package logic

import "imlogic/common/client"

func Close(c *client.Client) {
	service.hub.Unregister <- c
}
