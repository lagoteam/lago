package lago

import "github.com/goravel/framework/facades"

const Version string = "1.0.0"

type BootstrapServiceProvider struct {
}

func (r BootstrapServiceProvider) Boot() {
	//facades.Log.Debugf("[%#v] BootstrapServiceProvider.Boot()", "lago")
}

func (r BootstrapServiceProvider) Register() {
	//facades.Log.Debugf("[%#v] BootstrapServiceProvider.Register()", "lago")
}

func (r BootstrapServiceProvider) Install() {
	if facades.Config.GetBool("app.debug") {
		facades.Log.Debugf("[%#v] BootstrapServiceProvider.Install()", "lago")
	}
}

func (r BootstrapServiceProvider) Uninstall() {
}
