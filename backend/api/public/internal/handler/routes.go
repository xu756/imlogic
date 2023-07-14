// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/xu756/imlogic/api/public/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/get_code",
				Handler: GetCaptchaHandler(serverCtx),
			},
		},
		rest.WithPrefix("/imlogic/public/get_captcha"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/by_password",
				Handler: ByPasswordHandler(serverCtx),
			},
		},
		rest.WithPrefix("/imlogic/public/login"),
	)
}