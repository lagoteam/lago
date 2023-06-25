package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/modules/lago/http/controllers"
	"goravel/modules/lago/http/controllers/api/v1"
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
			indexController := controllers.NewIndexController()
			moduleRoute.Get("", indexController.Index)
		})

		apiRoute.Prefix("v1").Group(func(v1ApiRoute route.Route) {
			userController := v1.NewUserController()
			v1ApiRoute.Get("user/list", userController.List)
			v1ApiRoute.Post("user/add", userController.Add)
			v1ApiRoute.Get("user/detail/{id}", userController.Show)
		})

	})
}
