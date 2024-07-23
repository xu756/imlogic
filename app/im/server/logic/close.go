package logic

import (
	"imlogic/common/client"
	"log"
)

func onClose(c *client.Client) {
	service.hub.DelOneClient(c)
	log.Print("onclose")
}
