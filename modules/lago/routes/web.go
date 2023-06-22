package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	moduleControllers "goravel/modules/lago/http/controllers"
)

func Web() {
	facades.Route().Prefix("").Group(func(webRoute route.Route) {
		facades.Route().Prefix(facades.Config().GetString("lago.route_prefix")).Group(func(moduleRoute route.Route) {
			indexController := moduleControllers.NewIndexController()
			moduleRoute.Get("", indexController.Index)
		})
	})
}
