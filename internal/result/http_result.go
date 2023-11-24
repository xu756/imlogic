package result

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xu756/imlogic/internal/xerr"
	"net/http"
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
