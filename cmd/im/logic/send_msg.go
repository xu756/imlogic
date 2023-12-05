package logic

import (
	"github.com/xu756/imlogic/cmd/im/handler"
	"log"
)

func Imlogic(msg handler.Message) {
	log.Print(msg)
}
