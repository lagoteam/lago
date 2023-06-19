package providers

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http"
	"goravel/modules/example/routes"
)

type RouteServiceProvider struct {
}

func (r RouteServiceProvider) Boot() {

}

func (r RouteServiceProvider) Register() {
	//Add HTTP middlewares
	facades.Route.GlobalMiddleware(http.Kernel{}.Middleware()...)
	//Add routes
	routes.Api()
	routes.Web()
}
