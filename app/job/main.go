package main

import (
	"flag"

	"imlogic/app/job/logic"
	"imlogic/common/config"

	"github.com/cloudwego/kitex/pkg/klog"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	klog.SetLevel(klog.LevelFatal)
	logic.StartJob()
}
