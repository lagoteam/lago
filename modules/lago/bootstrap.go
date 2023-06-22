package lago

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	moduleProviders "goravel/modules/lago/providers"
)

const ModuleName string = "lago"
const ModuleVersion string = "1.0.0"
const ModulePrefixPath string = "/lago"

type BootstrapServiceProvider struct {
}

func (r BootstrapServiceProvider) Boot(app foundation.Application) {
	//facades.Log().Debugf("[%#v] BootstrapServiceProvider.Boot()", "lago")
	moduleProviders.AppServiceProvider{}.Boot(app)
}

func (r BootstrapServiceProvider) Register(app foundation.Application) {
	//facades.Log().Debugf("[%#v] BootstrapServiceProvider.Register()", "lago")
	moduleProviders.AppServiceProvider{}.Register(app)
}

func (r BootstrapServiceProvider) Install() {
	if facades.Config().GetBool("app.debug") {
		facades.Log().Debugf("[%#v] BootstrapServiceProvider.Install()", "lago")
	}
}

func (r BootstrapServiceProvider) Uninstall() {
}
