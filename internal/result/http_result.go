package result

import (
	"errors"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		var reserr *xerr.CodeError
		errors.As(err, &reserr)
		httpx.WriteJson(w, http.StatusUnauthorized, Error(reserr.GetCode(), reserr.GetMsg()))
	}
}
