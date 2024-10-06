package router

import (
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"imlogic/app/api/logic"
	"imlogic/common/config"
	"imlogic/internal/middleware"

	"github.com/cloudwego/hertz/pkg/app/server"
)

var HttpServer *server.Hertz

func InitRouter() {
	tracer, cfg := tracing.NewServerTracer()
	h := server.Default(tracer,
		server.WithHostPorts(config.RunData().Addr.ApiAddr),
		server.WithReadBufferSize(1024*1024*100),
		server.WithMaxRequestBodySize(1024*1024*100),
	)
	h.Use(tracing.ServerMiddleware(cfg))
	router := h.Group("/api")
	// 登录路由
	logic.LoginRoute(router.Group("/login"))
	// 用户路由
	userRouter := router.Group("/user")
	logic.UserRouter(userRouter)
	userRouter.Use(middleware.HertzJwt())
	// 会话路由
	chatRouter := router.Group("/chat")
	logic.ChatRoute(chatRouter)
	chatRouter.Use(middleware.HertzJwt())
	// 朋友路由
	friendRouter := router.Group("/friend")
	logic.FriendRoute(friendRouter)
	friendRouter.Use(middleware.HertzJwt())
	HttpServer = h
}
