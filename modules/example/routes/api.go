package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	moduleControllers "goravel/modules/example/http/controllers"
)

func Api() {
	facades.Route().Prefix("api").Group(func(apiRoute route.Route) {
		indexController := moduleControllers.NewIndexController()
		apiRoute.Get("/example", indexController.Index)
	})
}
