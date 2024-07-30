package main

import (
	"flag"

	"imlogic/common/config"
	"imlogic/common/mq"

	"github.com/cloudwego/kitex/pkg/klog"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	klog.SetLevel(klog.LevelFatal)
	rabbitmq := mq.NewRabbitMQPubSub("pubsub")
	defer rabbitmq.Destory()
	rabbitmq.RecieveSub()

}
