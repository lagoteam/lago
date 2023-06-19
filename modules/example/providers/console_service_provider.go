package providers

import (
	"github.com/goravel/framework/facades"
	"goravel/modules/example/console"
)

type ConsoleServiceProvider struct {
}

func (r ConsoleServiceProvider) Boot() {

}

func (r ConsoleServiceProvider) Register() {
	kernel := console.Kernel{}
	facades.Schedule.Register(kernel.Schedule())
	facades.Artisan.Register(kernel.Commands())
}
