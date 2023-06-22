package example

import (
	"github.com/goravel/framework/contracts/foundation"

	moduleProviders "goravel/modules/example/providers"
)

type ServiceProvider struct {
}

func (r ServiceProvider) Boot(app foundation.Application) {
	//facades.Log().Debugf("[%#v] ServiceProvider.Boot()", "example")
	moduleProviders.AppServiceProvider{}.Boot(app)
}

func (r ServiceProvider) Register(app foundation.Application) {
	//facades.Log().Debugf("[%#v] ServiceProvider.Register()", "example")
	moduleProviders.AppServiceProvider{}.Register(app)
}

/*func (r ServiceProvider) Install() {
	if facades.Config().GetBool("app.debug") {
		facades.Log().Debugf("[%#v] ServiceProvider.Install()", "lago")
	}
}

func (r ServiceProvider) Uninstall() {
}*/
