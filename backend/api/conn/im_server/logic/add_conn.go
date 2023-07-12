package logic

import (
	"context"
	"net/http"

	"github.com/xu756/imlogic/internal/pb"

	"github.com/thinkeridea/go-extend/exnet"
	"github.com/xu756/imlogic/api/conn/im_server/svc"
	"github.com/xu756/imlogic/app/chat/chat"
	"github.com/xu756/imlogic/internal/tool"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/xu756/imlogic/internal/xjwt"
	"github.com/zeromicro/go-zero/core/logx"
	"nhooyr.io/websocket"
)

var HostName string
var HeartBeat int
var chatRpc chat.Chat

type AddConnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConnLogic {
	return &AddConnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConnLogic) AddConn(w http.ResponseWriter, r *http.Request, userInfo *xjwt.AuthInfo) error {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	clientLogic := NewClientLogic(l.ctx, l.svcCtx)
	connId := tool.NewUid()
	ctx, cancelFunc := context.WithCancel(context.Background())
	client := &Client{
		connId:     connId,
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
	ws, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:         nil,
		InsecureSkipVerify:   false,
		OriginPatterns:       nil,
		CompressionMode:      0,
		CompressionThreshold: 0,
	})
	if err != nil {
		return xerr.NewSystemError("【websocket.Accept】升级请求头 error:" + err.Error())
	}
	err = client.setWs(ws)
	if err != nil {
		return xerr.NewSystemError("【websocket.Accept】设置ws" + err.Error())
	}
	Hubs.register <- client
	clientLogic.client = client
	chatRpc = l.svcCtx.Chat
	chatRpc.Meta(l.ctx, &pb.ImMeta{
		DetailType: 0,
		ConnId:     connId,
		UserId:     userInfo.ID,
		Role:       userInfo.Role,
		Ip:         ip,
		Issue:      userInfo.Issuer,
		Version:    "v1",
		ImServer:   HostName,
		Data:       nil,
	})
	return nil
}
