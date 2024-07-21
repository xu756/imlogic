package handler

import (
	"imlogic/common/config"
	"testing"
)

func TestInitRouter(t *testing.T) {
	config.Init("../../../../configs/dev.yaml")
	go hub.Run()
	InitRouter()
	HttpServer.Spin()
}
