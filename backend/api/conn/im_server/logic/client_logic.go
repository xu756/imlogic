package logic

import (
	"context"
	"log"

	"github.com/xu756/imlogic/api/conn/im_server/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"nhooyr.io/websocket"
)

type ClientLogic struct {
	logx.Logger
	ctx    context.Context
	client *Client
	svcCtx *svc.ServiceContext
}

func NewClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientLogic {
	return &ClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *Client) readMessage() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			ws := c.getWs()
			if ws == nil {
				return
			}
			typ, message, err := ws.Read(c.ctx)
			if err != nil {
				c.Close(websocket.StatusInternalError, "read message error")
				return
			}
			log.Print(typ, message)
			if typ == websocket.MessageBinary {
				go c.msgLogic(message)
			}
		}
	}
}

func (c *Client) WriteMessage(msg []byte) error {
	ws := c.getWs()
	if ws == nil {
		return nil
	}
	return ws.Write(c.ctx, websocket.MessageBinary, msg)

}
