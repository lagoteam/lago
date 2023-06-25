package services

import (
	contractsHttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/log"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http"
)

type BaseService struct {
	Ctx contractsHttp.Context
	Log log.Log
}

func NewBaseService() BaseService {
	return BaseService{
		Ctx: http.Background(),
		Log: facades.Log(),
	}
}
