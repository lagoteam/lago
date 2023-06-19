package example

import "goravel/modules/example/providers"

type BootstrapServiceProvider struct {
}

func (r BootstrapServiceProvider) Boot() {
	providers.AppServiceProvider{}.Boot()
}

func (r BootstrapServiceProvider) Register() {
	providers.AppServiceProvider{}.Register()
}
