package commands

import (
	"github.com/gookit/color"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"goravel/database/migrations"
	"goravel/packages/migrate"
)

type MigrateUpCommand struct {
	config config.Config
}

func NewMigrateUpCommand(config config.Config) *MigrateUpCommand {
	return &MigrateUpCommand{
		config: config,
	}
}

// Signature The name and signature of the console command.
func (receiver *MigrateUpCommand) Signature() string {
	return "gorm:migrate"
}

// Description The console command description.
func (receiver *MigrateUpCommand) Description() string {
	return "Run the gorm migrations"
}

// Extend The console command extend.
func (receiver *MigrateUpCommand) Extend() command.Extend {
	return command.Extend{
		Category: "gorm",
	}
}

// Handle Execute the console command.
func (receiver *MigrateUpCommand) Handle(ctx console.Context) error {
	// 注册 database/migrations 下的所有迁移文件
	migrations.Initialize()

	// 执行迁移
	migrationManager, err := migrate.NewMigrationManager(receiver.config)
	if err != nil {
		return err
	}

	color.Yellowln("Migration start up")
	migrationManager.Up()
	color.Greenln("Migration success up")

	return nil
}
