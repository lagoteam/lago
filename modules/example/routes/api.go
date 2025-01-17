package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	moduleControllers "goravel/modules/example/http/controllers"
)

func Api() {
	facades.Route().Prefix("api").Group(func(apiRoute route.Route) {
		apiRoute.Prefix(facades.Config().GetString("example.route_prefix")).Group(func(moduleRoute route.Route) {
			moduleRoute.Get("ping", func(ctx http.Context) {
				ctx.Response().Json(http.StatusOK, http.Json{
					"name":    facades.Config().GetString("app.name"),
					"version": facades.Config().GetString("app.version"),
				})
			})
			indexController := moduleControllers.NewIndexController()
			moduleRoute.Get("", indexController.Index)
		})
	})
}
