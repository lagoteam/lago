package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	moduleControllers "goravel/modules/lago/http/controllers"
)

func Api() {
	facades.Route().Prefix("api").Group(func(apiRoute route.Route) {
		indexController := moduleControllers.NewIndexController()
		apiRoute.Get("/lago", indexController.Index)
	})
}
