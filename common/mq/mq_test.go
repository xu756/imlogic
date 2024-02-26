package mq

import (
	"server/common/config"
	"testing"
)

//import "server/common/config"
//
//func ()  {
//	config.Init("../../configs/dev.yaml")
//
//}

func TestInit(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	Init()

}
