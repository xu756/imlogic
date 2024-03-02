package handler

import (
	"github.com/xu756/imlogic/kitex_gen/im"
)

func (i ImRpcImpl) Receive(stream im.ImSrv_ReceiveServer) (err error) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		err = stream.Send(&im.Message{
			MsgId:     msg.MsgId,
			From:      msg.From,
			To:        msg.To,
			Timestamp: msg.Timestamp,
			Params: map[string]string{
				"send": "ok",
			},
		})
		if err != nil {
			stream.Close()
			return err
		}
	}
}
