package console

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	events := []schedule.Event{}
	return events
}

func (kernel *Kernel) Commands() []console.Command {
	commands := []console.Command{}
	return commands
}
