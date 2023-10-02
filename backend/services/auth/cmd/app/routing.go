// Пакет роутов
// BaseRoutes - перечень всех маршрутов не содержащие middleware,
// содержащий handler,привязанной фичи и всех middleware для handler
package app

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/middleware"
	"github.com/PestovOleg/mini-bank/backend/services/auth/internal/http/handler/v1/auth"
	"github.com/PestovOleg/mini-bank/backend/services/auth/internal/http/handler/v1/health"
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	Handler     http.Handler
	Feature     string
	Middlewares []mux.MiddlewareFunc
}

func BaseRoutes(s *Services) map[string]map[string]RouteConfig {
	return map[string]map[string]RouteConfig{
		"/health": {
			http.MethodGet: {
				Handler: health.NewHealthCheckHandler(),
			},
		},
		"/auth": {
			http.MethodPost: {
				Handler:     auth.NewAuthHandler(s.AuthService).CreateAuth(),
				Feature:     "CreateUserToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware},
			},
			http.MethodGet: {
				Handler:     auth.NewAuthHandler(s.AuthService).Authenticate(),
				Feature:     "AuthenticateToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware},
			},
		},
		"/auth/{id}": {
			http.MethodGet: {
				Handler:     auth.NewAuthHandler(s.AuthService).Authorize(),
				Feature:     "AuthorizeToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware},
			},
			http.MethodDelete: {
				Handler:     auth.NewAuthHandler(s.AuthService).DeleteAuth(),
				Feature:     "DeleteUserToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware, middleware.BasicAuthMiddleware},
			},
		},
	}
}

func SetHandler(r *mux.Router, paths map[string]map[string]RouteConfig) {
	for path, methods := range paths {
		for method, config := range methods {
			handler := config.Handler
			for _, middleware := range config.Middlewares { //оборачиваем во все middleware
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
