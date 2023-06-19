package example

import (
	"github.com/goravel/framework/contracts/foundation"
	"goravel/modules/example/providers"
)

type BootstrapServiceProvider struct {
}

func (r BootstrapServiceProvider) Boot(app foundation.Application) {
	providers.AppServiceProvider{}.Boot(app)
}

func (r BootstrapServiceProvider) Register(app foundation.Application) {
	providers.AppServiceProvider{}.Register(app)
}
