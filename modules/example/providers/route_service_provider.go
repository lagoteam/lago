package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	"goravel/app/http"
	"goravel/modules/example/routes"
)

type RouteServiceProvider struct {
}

func (r RouteServiceProvider) Boot(app foundation.Application) {

}

func (r RouteServiceProvider) Register(app foundation.Application) {
	//Add HTTP middlewares
	facades.Route().GlobalMiddleware(http.Kernel{}.Middleware()...)
	//Add routes
	routes.Api()
	routes.Web()
}
