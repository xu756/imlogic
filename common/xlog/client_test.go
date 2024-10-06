package xlog

import (
	"testing"
	"time"
)

func TestSetUp(t *testing.T) {

	for i := 0; i < 100; i++ {
		Send(map[string]interface{}{
			"_timestamp": time.Now().UnixMicro(),
			"name":       "xu756",
			"msg":        "你好",
			"log":        "ip-10-2-56-221.us-east-2.compute.internal",
			"i":          i,
		})
	}
	time.Sleep(time.Second * 10)
	for i := 0; i < 100; i++ {
		Send(map[string]interface{}{
			"_timestamp": time.Now().UnixMicro(),
			"name":       "xu756",
			"msg":        "你好",
			"log":        "ip-10-2-56-221.us-east-2.compute.internal",
			"i":          i,
		})
	}
	time.Sleep(time.Second * 10)

}
