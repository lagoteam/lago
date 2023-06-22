package middleware

import (
	"github.com/gin-contrib/static"

	contractsHttp "github.com/goravel/framework/contracts/http"
	frameworkHttp "github.com/goravel/framework/http"
)

func Static() contractsHttp.Middleware {
	return func(ctx contractsHttp.Context) {
		static.Serve("/", static.LocalFile("./public", false))(ctx.(*frameworkHttp.GinContext).Instance())
		ctx.Request().Next()
	}
}
