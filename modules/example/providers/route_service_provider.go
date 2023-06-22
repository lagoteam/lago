package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	moduleHttp "goravel/modules/example/http"
	moduleRoutes "goravel/modules/example/routes"
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
