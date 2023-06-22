package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type IndexController struct {
	//Dependent services
}

func NewIndexController() *IndexController {
	return &IndexController{
		//Inject services
	}
}

func (r *IndexController) Index(ctx http.Context) {
	ctx.Response().Json(http.StatusOK, http.Json{
		"name":           facades.Config().GetString("app.name"),
		"version":        facades.Config().GetString("app.version"),
		"module_name":    facades.Config().GetString("lago.name"),
		"module_version": facades.Config().GetString("lago.version"),
	})
}
