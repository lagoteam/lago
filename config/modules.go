package config

import (
	"github.com/goravel/framework/contracts"
	"github.com/goravel/framework/facades"

	"goravel/modules/example"
	"goravel/modules/lago"
)

func init() {
	config := facades.Config
	config.Add("modules", map[string]interface{}{
		"enabled": config.Env("MODULES_ENABLED", true),

		"providers": []contracts.ServiceProvider{
			&example.BootstrapServiceProvider{},
			&lago.BootstrapServiceProvider{},
		},

		"paths": map[string]interface{}{
			"modules":   config.Env("MODULES_PATHS_MODULES", "modules"),
			"assets":    config.Env("MODULES_PATHS_ASSETS", "public"),
			"migration": config.Env("MODULES_PATHS_MIGRATION", "database/migrations"),
		},
	})
}
