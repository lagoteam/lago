package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	moduleHttp "goravel/modules/lago/http"
	moduleRoutes "goravel/modules/lago/routes"
	"goravel/modules/lago/utils/result"
)

type RouteServiceProvider struct {
}

func (r RouteServiceProvider) Boot(app foundation.Application) {

}

func (r RouteServiceProvider) Register(app foundation.Application) {
	//Add HTTP middlewares
	facades.Route().GlobalMiddleware(moduleHttp.Kernel{}.Middleware()...)
	//Add HTTP 404 Error
	facades.Route().Fallback(func(ctx http.Context) {
		//ctx.Response().String(404, "not found")
		res := result.Fail().SetCode(result.CodeNotFound).SetMessage(result.CodeMessage[result.CodeNotFound])
		ctx.Response().Json(200, res)
	})
	//Add routes
	moduleRoutes.Api()
	moduleRoutes.Web()
}
