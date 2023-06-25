package console

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
	"github.com/goravel/framework/facades"
	"goravel/modules/lago/console/commands"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	scheduleEvents := []schedule.Event{}
	return scheduleEvents
}

func (kernel *Kernel) Commands() []console.Command {
	config := facades.Config()
	consoleCommands := []console.Command{}
	consoleCommands = append(consoleCommands, &commands.SendEmailCommand{})
	consoleCommands = append(consoleCommands, commands.NewMigrateUpCommand(config))
	consoleCommands = append(consoleCommands, commands.NewMigrateRollbackCommand(config))
	return consoleCommands
}
