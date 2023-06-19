package providers

import (
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type EventServiceProvider struct {
}

func (r EventServiceProvider) Boot(app foundation.Application) {

}

func (r EventServiceProvider) Register(app foundation.Application) {
	facades.Event().Register(r.Listen())
}

func (r EventServiceProvider) Listen() map[event.Event][]event.Listener {
	return map[event.Event][]event.Listener{}
}
