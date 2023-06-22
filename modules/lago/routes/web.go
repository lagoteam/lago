package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Web() {
	facades.Route().Prefix("lago").Group(func(webRoute route.Route) {
		webRoute.Get("", func(ctx http.Context) {
			ctx.Response().Json(http.StatusOK, http.Json{
				"name":    facades.Config().GetString("app.name"),
				"version": facades.Config().GetString("app.version"),
			})
		})
	})
}
