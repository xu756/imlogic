package main

import (
	"imlogic/app/job/logic"

	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	klog.SetLevel(klog.LevelFatal)
	logic.InitService()
	logic.StartJob()
}
