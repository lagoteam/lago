package example

import (
	"github.com/goravel/framework/contracts/foundation"

	moduleProviders "goravel/modules/example/providers"
)

const ModuleName string = "example"
const ModuleVersion string = "1.0.0"

//const ModulePrefixPath string = "/example"

type BootstrapServiceProvider struct {
}

func (r BootstrapServiceProvider) Boot(app foundation.Application) {
	moduleProviders.AppServiceProvider{}.Boot(app)
}

func (r BootstrapServiceProvider) Register(app foundation.Application) {
	moduleProviders.AppServiceProvider{}.Register(app)
}
