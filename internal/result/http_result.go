package result

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/xu756/imlogic/internal/xerr"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func HttpSuccess(ctx *app.RequestContext, resp interface{}) {
	r := successRes(resp)
	ctx.JSON(http.StatusOK, r)
	ctx.Abort()
}

func HttpError(ctx *app.RequestContext, err error) {
	hlog.Debug(err.(remote.TransError).Unwrap().(xerr.CodeError))
	var resErr xerr.CodeError
	errors.As(errors.Unwrap(err), &resErr)
	ctx.JSON(http.StatusOK, errorRes(resErr.Code, resErr.Msg))
	ctx.Abort()

}
