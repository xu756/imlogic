package middleware

import (
	"context"
	"imlogic/internal/result"
	"imlogic/internal/xerr"

	"github.com/cloudwego/hertz/pkg/app"
)

func HertzJwt() app.HandlerFunc {
	// todo 验证jwt
	return func(ctx context.Context, c *app.RequestContext) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			result.HttpError(c, xerr.RoleErr("Authorization is empty"))
			return
		}

	}

}
