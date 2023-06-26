package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type SendEmailCommand struct {
}

// Signature The name and signature of the console command.
func (receiver *SendEmailCommand) Signature() string {
	return "send:email"
}

// Description The console command description.
func (receiver *SendEmailCommand) Description() string {
	return "Send email"
}

// Extend The console command extend.
func (receiver *SendEmailCommand) Extend() command.Extend {
	return command.Extend{
		Category: "email",
	}
}

// Handle Execute the console command.
func (receiver *SendEmailCommand) Handle(ctx console.Context) error {
	return nil
}
