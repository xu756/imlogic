package handler

import (
	"github.com/xu756/imlogic/common/config"
	"testing"
)

func TestInitRouter(t *testing.T) {
	config.Init("../../../../configs/dev.yaml")
	go Hub.Run()
	InitRouter()
	HttpServer.Spin()
}
