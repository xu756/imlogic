package handler

import (
	"github.com/xu756/imlogic/common/db"
	"github.com/xu756/imlogic/kitex_gen/im"
	"log"
)

type ImRpcImpl struct {
	Model db.Model
}

func NewImRpcImpl() *ImRpcImpl {
	return &ImRpcImpl{
		Model: db.NewModel(),
	}
}

func (i ImRpcImpl) Receive(stream im.ImSrv_ReceiveServer) (err error) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Print("receive msg: ", msg)
		log.Print(msg.Parms["a"])
	}
}
