package commands

import (
	"github.com/gookit/color"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"goravel/database/migrations"
	"goravel/packages/migrate"
)

type MigrateRollbackCommand struct {
	config config.Config
}

func NewMigrateRollbackCommand(config config.Config) *MigrateRollbackCommand {
	return &MigrateRollbackCommand{
		config: config,
	}
}

// Signature The name and signature of the console command.
func (receiver *MigrateRollbackCommand) Signature() string {
	return "gorm:migrate-rollback"
}

// Description The console command description.
func (receiver *MigrateRollbackCommand) Description() string {
	return "Run the gorm migrations"
}

// Extend The console command extend.
func (receiver *MigrateRollbackCommand) Extend() command.Extend {
	return command.Extend{
		Category: "gorm",
	}
}

// Handle Execute the console command.
func (receiver *MigrateRollbackCommand) Handle(ctx console.Context) error {
	// 注册 database/migrations 下的所有迁移文件
	migrations.Initialize()

	// 执行迁移
	migrationManager, err := migrate.NewMigrationManager(receiver.config)
	if err != nil {
		return err
	}

	color.Yellowln("Migration start rollback")
	migrationManager.Rollback()
	color.Greenln("Migration success rollback")

	return nil
}
