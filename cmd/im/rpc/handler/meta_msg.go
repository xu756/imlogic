package handler

import (
	"context"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/kitex_gen/im"
	"log"
)

func (i ImRpcImpl) MetaMsg(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	switch req.MsgMeta.DetailType {
	case "connect":
		// todo
		err := i.Cache.AddConn(ctx, req.Params["userId"], req.Device, req.From, req.Params["hostName"], req.Timestamp)
		if err != nil {
			return nil, xerr.CacheErr()
		}
		break
	case "disconnect":
		// todo
		log.Printf("%s[%s] disconnect ", req.Params["userId"], req.From)
		break
	case "heartbeat":
		// todo
		log.Printf("%s[%s] heartbeat ", req.Params["userId"], req.From)
		break
	default:
		// todo
		log.Printf("%s[%s] unknown ", req.Params["userId"], req.From)
		break

	}
	return &im.MessageRes{
		MsgId:   req.MsgId,
		From:    req.To,
		To:      req.From,
		Success: true,
	}, nil
}
