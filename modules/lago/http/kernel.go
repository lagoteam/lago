package http

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/http/middleware"

	moduleMiddleware "goravel/modules/lago/http/middleware"
)

type Kernel struct {
}

// Middleware The application's global HTTP middleware stack.
// These middleware are run during every request to your application.
func (kernel Kernel) Middleware() []http.Middleware {
	middlewares := []http.Middleware{
		middleware.Cors(),
		moduleMiddleware.Static(),
	}
	return middlewares
}
