package handler

import (
	"net/http"

	"github.com/xu756/imlogic/api/conn/im_server/svc"
)

func RegisterHandlers(server *http.Server, serverCtx *svc.ServiceContext) {
	// 设置路由、中间件、处理新连接的函数
	mux := http.NewServeMux()

	mux.HandleFunc("/imlogic/conn", addConn(serverCtx))
	server.Handler = mux
}
