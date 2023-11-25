package result

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xu756/imlogic/internal/xerr"
	"net/http"
)

func HttpSuccess(ctx *app.RequestContext, resp interface{}) {
	r := successRes(resp)
	ctx.JSON(http.StatusOK, r)
	ctx.Abort()
}

func HttpError(ctx *app.RequestContext, err error) {
	var resErr xerr.CodeError
	errors.As(err, &resErr)
	ctx.JSON(http.StatusOK, errorRes(resErr.Code, resErr.Msg))
	ctx.Abort()
}
