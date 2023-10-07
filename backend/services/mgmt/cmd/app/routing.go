// Пакет роутов версии 1
// BaseRoutes - маршруты не содержащие middleware
//
//	[L] - включающие middleware: LoggingMiddleware
//	[LA]- включающие middleware: LoggingMiddleware,AuthorizeMiddleware
package app

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/middleware"
	"github.com/PestovOleg/mini-bank/backend/services/mgmt/internal/http/handler/v1/health"
	handlerMgmt "github.com/PestovOleg/mini-bank/backend/services/mgmt/internal/http/handler/v1/mgmt"
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	Handler     http.Handler
	Feature     string
	Middlewares []mux.MiddlewareFunc
}

func BaseRoutes(s *Services) map[string]map[string]RouteConfig {
	return map[string]map[string]RouteConfig{
		"/mgmt-minibank-health": {
			http.MethodGet: {
				Handler: health.NewHealthCheckHandler(),
			},
		},
		"/mgmt": {
			http.MethodPost: {
				Handler:     handlerMgmt.NewMgmtHandler().CreateUser(),
				Feature:     "CreateUserToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware},
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
