package config

import (
	"net/http"

	"github.com/xu756/imlogic/internal/result"

	"github.com/xu756/imlogic/api/public/internal/logic/config"
	"github.com/xu756/imlogic/api/public/internal/svc"
)

func GetConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := config.NewGetConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetConfig()
		result.HttpResult(r, w, resp, err)
	}
}
