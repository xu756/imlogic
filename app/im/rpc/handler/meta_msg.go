package handler

import (
	"context"
	"imlogic/internal/xerr"
	"imlogic/kitex_gen/im"
)

func (i ImRpcImpl) MetaMsg(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	res = &im.MessageRes{}
	switch req.MsgMeta.DetailType {
	case "connect":
		// todo
		err := i.Cache.AddConn(ctx, req.From, req.Device, req.LinkId, req.Hostname, req.Timestamp)
		if err != nil {
			return nil, xerr.CacheErr()
		}
		break
	case "disconnect":
		// todo
		//log.Printf("%s[%s] disconnect ", req.From, req.LinkId)
		break
	case "heartbeat":
		// todo
		//log.Printf("%s[%s] heartbeat ", req.From, req.LinkId)
		break
	default:
		// todo
		//log.Printf("%s[%s] unknown ", req.From, req.LinkId)
		break

	}
	res.MsgId = req.MsgId
	res.From = req.To
	res.To = req.From
	res.Success = true
	return
}
