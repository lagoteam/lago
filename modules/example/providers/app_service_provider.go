package providers

import "github.com/goravel/framework/contracts/foundation"

type AppServiceProvider struct {
}

func (r AppServiceProvider) Boot(app foundation.Application) {
	ConfigServiceProvider{}.Boot(app)
	RouteServiceProvider{}.Boot(app)
	ConsoleServiceProvider{}.Boot(app)
	EventServiceProvider{}.Boot(app)
	QueueServiceProvider{}.Boot(app)
}

func (r AppServiceProvider) Register(app foundation.Application) {
	ConfigServiceProvider{}.Register(app)
	RouteServiceProvider{}.Register(app)
	ConsoleServiceProvider{}.Register(app)
	EventServiceProvider{}.Register(app)
	QueueServiceProvider{}.Register(app)
}
