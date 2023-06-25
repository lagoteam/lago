package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	moduleControllers "goravel/modules/lago/http/controllers"
)

func Api() {
	facades.Route().Prefix("api").Group(func(apiRoute route.Route) {
		apiRoute.Prefix(facades.Config().GetString("lago.route_prefix")).Group(func(moduleRoute route.Route) {
			moduleRoute.Get("ping", func(ctx http.Context) {
				ctx.Response().Json(http.StatusOK, http.Json{
					"name":    facades.Config().GetString("app.name"),
					"version": facades.Config().GetString("app.version"),
				})
			})
			indexController := moduleControllers.NewIndexController()
			moduleRoute.Get("", indexController.Index)

			userController := moduleControllers.NewUserController()
			moduleRoute.Get("user/list", userController.List)
			moduleRoute.Post("user/add", userController.Add)
			moduleRoute.Get("user/detail/{id}", userController.Show)
		})
	})
}
