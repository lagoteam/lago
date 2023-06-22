package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	moduleHttp "goravel/modules/lago/http"
	moduleRoutes "goravel/modules/lago/routes"
)

type RouteServiceProvider struct {
}

func (r RouteServiceProvider) Boot(app foundation.Application) {

}

func (r RouteServiceProvider) Register(app foundation.Application) {
	//Add HTTP middlewares
	facades.Route().GlobalMiddleware(moduleHttp.Kernel{}.Middleware()...)
	//Add routes
	moduleRoutes.Api()
	moduleRoutes.Web()
}
