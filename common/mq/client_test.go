package mq

import (
	"fmt"
	"imlogic/common/config"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	log.Print(config.RunData.MqUrl)
	client := NewClient("im/message")
	defer client.Close()
	// 循环10次
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("hello %d", i)
		err := client.WorkPublish(msg, "group")
		if err != nil {
			log.Fatal(err)
		}
	}
}
