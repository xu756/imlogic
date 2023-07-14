package handler

import (
	"github.com/xu756/imlogic/internal/result"
	"net/http"

	"github.com/xu756/imlogic/api/public/internal/logic"
	"github.com/xu756/imlogic/api/public/internal/svc"
	"github.com/xu756/imlogic/api/public/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ByPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewByPasswordLogic(r.Context(), svcCtx)
		resp, err := l.ByPassword(&req)
		result.HttpResult(r, w, resp, err)
	}
}
