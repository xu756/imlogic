package result

import (
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
	errId := err.(*remote.TransError).TypeID()
	ctx.JSON(http.StatusOK, errorRes(errId, xerr.GetMsg(errId)))
	ctx.Abort()
}
