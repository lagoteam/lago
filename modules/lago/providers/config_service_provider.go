package providers

import (
	"github.com/goravel/framework/contracts/foundation"

	moduleConfig "goravel/modules/lago/config"
)

type ConfigServiceProvider struct {
}

func (r ConfigServiceProvider) Boot(app foundation.Application) {
}

func (r ConfigServiceProvider) Register(app foundation.Application) {
	moduleConfig.Boot()
}
