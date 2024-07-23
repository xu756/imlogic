package result

import (
	"errors"
	"imlogic/internal/xerr"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func HttpSuccess(c *app.RequestContext, resp interface{}) {
	r := successRes(resp)
	c.JSON(http.StatusOK, r)
	c.Abort()
}

func HttpError(c *app.RequestContext, err error) {
	var resErr xerr.CodeError
	errors.As(err, &resErr)
	c.JSON(http.StatusOK, errorRes(resErr.Code, resErr.Msg))
	c.Abort()
}

func HttpParamErr(c *app.RequestContext) {
	c.JSON(http.StatusOK, errorRes(xerr.ParamErrCode, xerr.GetMsg(xerr.ParamErrCode)))
	c.Abort()
}
