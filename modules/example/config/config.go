package config

import (
	"github.com/goravel/framework/facades"
)

func Boot() {}

func init() {
	config := facades.Config()
	config.Add("lago", map[string]any{
		"name":    config.Env("MODULE_EXAMPLE_NAME", "example"),
		"version": config.Env("MODULE_EXAMPLE_VERSION", "1.0.0"),
	})
}
