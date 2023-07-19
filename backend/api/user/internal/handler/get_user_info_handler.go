package handler

import (
	"net/http"

	"github.com/xu756/imlogic/api/user/internal/logic"
	"github.com/xu756/imlogic/api/user/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		result.HttpResult(r, w, resp, err)
	}
}
