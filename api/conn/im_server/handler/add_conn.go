package handler

import (
	"net/http"

	"github.com/xu756/imlogic/api/conn/im_server/logic"
	"github.com/xu756/imlogic/api/conn/im_server/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

func addConn(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo, err := svcCtx.Jwt.GetTokenFromHeader(r, "Authorization")
		if err != nil {
			// todo 未登录
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if _, err = svcCtx.RedisClient.Get("cache:appserver:user:id:" + userInfo.GetStrId()); err != nil {
			logx.WithContext(r.Context()).Error("中间件 middleware error:", err)
			return
		}

		l := logic.NewAddConnLogic(r.Context(), svcCtx)
		err = l.AddConn(w, r, userInfo)
		if err != nil {
			logx.WithContext(r.Context()).Error("logic error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
