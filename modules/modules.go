package modules

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type BootstrapServiceProvider struct {
}

func (r BootstrapServiceProvider) Boot(app foundation.Application) {
	serviceProviders := facades.Config().Get("modules.providers").([]foundation.ServiceProvider)
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Boot(app)
	}
}

func (r BootstrapServiceProvider) Register(app foundation.Application) {
	serviceProviders := facades.Config().Get("modules.providers").([]foundation.ServiceProvider)
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Register(app)
	}
}
