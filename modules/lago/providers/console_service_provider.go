package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	"goravel/modules/lago/console"
)

type ConsoleServiceProvider struct {
}

func (r ConsoleServiceProvider) Boot(app foundation.Application) {

}

func (r ConsoleServiceProvider) Register(app foundation.Application) {
	kernel := console.Kernel{}
	facades.Schedule().Register(kernel.Schedule())
	facades.Artisan().Register(kernel.Commands())
}
