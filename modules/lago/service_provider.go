package lago

import (
	"github.com/goravel/framework/contracts/foundation"

	moduleProviders "goravel/modules/lago/providers"
)

type ServiceProvider struct {
}

func (r ServiceProvider) Boot(app foundation.Application) {
	moduleProviders.AppServiceProvider{}.Boot(app)
}

func (r ServiceProvider) Register(app foundation.Application) {
	moduleProviders.AppServiceProvider{}.Register(app)
}
