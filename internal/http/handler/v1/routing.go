// Пакет роутов версии 1
// BaseRoutes - маршруты не содержащие middleware
//
//	[L] - включающие middleware: LoggingMiddleware
//	[LA]- включающие middleware: LoggingMiddleware,AuthorizeMiddleware
package v1

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/internal/common"
	"github.com/PestovOleg/mini-bank/internal/http/handler/v1/health"
	handlerUser "github.com/PestovOleg/mini-bank/internal/http/handler/v1/user"
	"github.com/gorilla/mux"
)

func BaseRoutes(s *common.Services) map[string]map[string]http.Handler {
	return map[string]map[string]http.Handler{
		"/health": {http.MethodGet: health.NewHealthCheckHandler()},
		"/users":  {http.MethodPost: handlerUser.NewUserHandler(s.UserService)},
	}
}

func BaseRoutesL(s *common.Services) map[string]map[string]http.Handler {
	return map[string]map[string]http.Handler{
		"/users": {http.MethodGet: handlerUser.NewUserHandler(s.UserService)},
	}
}

func SetHandler(r *mux.Router, paths map[string]map[string]http.Handler) {
	for path := range paths {
		for method, handler := range paths[path] {
			r.Handle(path, handler).Methods(method)
		}
	}
}
