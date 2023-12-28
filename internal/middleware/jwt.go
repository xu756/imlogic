package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func HertzJwt() app.HandlerFunc {
	// todo 验证jwt
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.Next(c)
	}

}
