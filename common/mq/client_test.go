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

	// 循环10次
	for i := 0; i < 1000; i++ {
		client := NewClient("im/message")
		msg := fmt.Sprintf("helllo %d", i)
		err := client.Publish(msg)
		if err != nil {
			log.Fatal(err)
		}
	}

}
