package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"imlogic/internal/result"
	"imlogic/internal/xerr"
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
