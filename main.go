package main

import (
	im "imlogic/kitex_gen/im/imserver"
	"log"
)

func main() {
	svr := im.NewServer(new(ImServerImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
