package modules

import (
	"github.com/goravel/framework/contracts"
	"github.com/goravel/framework/facades"
)

type BootstrapServiceProvider struct {
}

func (r BootstrapServiceProvider) Boot() {
	serviceProviders := facades.Config.Get("modules.providers").([]contracts.ServiceProvider)
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Boot()
	}
}

func (r BootstrapServiceProvider) Register() {
	serviceProviders := facades.Config.Get("modules.providers").([]contracts.ServiceProvider)
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Register()
	}
}
