package result

import (
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/xu756/imlogic/internal/xerr"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func HttpSuccess(c *app.RequestContext, resp interface{}) {
	r := successRes(resp)
	c.JSON(http.StatusOK, r)
	c.Abort()
}

func HttpError(c *app.RequestContext, err error) {
	errId := err.(*remote.TransError).TypeID()
	c.JSON(http.StatusOK, errorRes(errId, xerr.GetMsg(errId)))
	c.Abort()
}
