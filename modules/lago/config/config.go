package config

import (
	"github.com/goravel/framework/facades"
)

func Boot() {}

func init() {
	config := facades.Config()
	config.Add("lago", map[string]any{
		"name":    config.Env("MODULE_LAGO_NAME", "lago"),
		"version": config.Env("MODULE_LAGO_VERSION", "1.0.0"),
	})
}
