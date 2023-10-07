// Пакет роутов версии 1
// BaseRoutes - маршруты не содержащие middleware
//
//	[L] - включающие middleware: LoggingMiddleware
//	[LA]- включающие middleware: LoggingMiddleware,AuthorizeMiddleware
package app

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/middleware"
	"github.com/PestovOleg/mini-bank/backend/services/user/internal/http/handler/v1/health"
	handlerUser "github.com/PestovOleg/mini-bank/backend/services/user/internal/http/handler/v1/user"
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	Handler     http.Handler
	Feature     string
	Middlewares []mux.MiddlewareFunc
}

func BaseRoutes(s *Services) map[string]map[string]RouteConfig {
	return map[string]map[string]RouteConfig{
		"/user-minibank-health": {
			http.MethodGet: {
				Handler: health.NewHealthCheckHandler(),
			},
		},
		"/users": {
			http.MethodPost: {
				Handler:     handlerUser.NewUserHandler(s.UserService).CreateUser(),
				Feature:     "CreateUserToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware},
			},
		},
		"/users/{id}": {
			http.MethodGet: {
				Handler:     handlerUser.NewUserHandler(s.UserService).GetUser(),
				Feature:     "GetUserToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware, middleware.BasicAuthMiddleware},
			},
			http.MethodPut: {
				Handler:     handlerUser.NewUserHandler(s.UserService).UpdateUser(),
				Feature:     "UpdateUserToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware, middleware.BasicAuthMiddleware},
			},
		},
	}
}

func SetHandler(r *mux.Router, paths map[string]map[string]RouteConfig) {
	for path, methods := range paths {
		for method, config := range methods {
			handler := config.Handler
			for _, middleware := range config.Middlewares { // оборачиваем во все middleware
				handler = middleware(handler)
			}

			if config.Feature == "" { // присваиваем фичи
				r.Handle(path, handler).Methods(method)
			} else {
				r.Handle(path, middleware.FeatureToggleMiddleware(config.Feature, handler)).Methods(method)
			}
		}
	}
}
