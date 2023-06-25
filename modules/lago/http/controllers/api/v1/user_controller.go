package v1

import (
	"strconv"

	"github.com/goravel/framework/contracts/http"

	"goravel/modules/lago/models"
	"goravel/modules/lago/services"
	"goravel/modules/lago/utils/result"
)

type UserController struct {
	//Dependent services
	Service services.UserService
}

func NewUserController() *UserController {
	userService := services.NewUserService()
	return &UserController{
		//Inject services
		Service: userService,
	}
}

func (r *UserController) List(ctx http.Context) {
	page, _ := strconv.Atoi(ctx.Request().Input("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Input("limit", "10"))
	search := ctx.Request().Input("search")
	items, total, err := r.Service.List(page, limit, search)
	if err != nil {
		ctx.Response().Success().Json(result.Fail().SetMessage(err.Error()))
		return
	}
	data := map[string]interface{}{
		"items": items,
		"total": total,
	}
	ctx.Response().Success().Json(result.Success().SetData(data))
}

func (r *UserController) Add(ctx http.Context) {
	username := ctx.Request().Input("username")
	password := ctx.Request().Input("password")

	var user models.User
	user.Username = username
	user.Password = password

	user, err := r.Service.CreateUser(user)
	if err != nil {
		ctx.Response().Success().Json(result.Fail().SetMessage(err.Error()))
		return
	}
	ctx.Response().Success().Json(result.Success().SetData(user))
}

func (r *UserController) Show(ctx http.Context) {
	id, _ := strconv.ParseUint(ctx.Request().Input("id"), 0, 64)
	user, err := r.Service.GetUser(uint(id))
	if err != nil {
		ctx.Response().Success().Json(result.Fail().SetMessage(err.Error()))
		return
	}
	ctx.Response().Success().Json(result.Success().SetData(user))
}
