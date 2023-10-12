// Пакет роутов версии 1
// BaseRoutes - маршруты не содержащие middleware
//
//	[L] - включающие middleware: LoggingMiddleware
//	[LA]- включающие middleware: LoggingMiddleware,AuthorizeMiddleware
package app

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/middleware"
	"github.com/PestovOleg/mini-bank/backend/services/uproxy/internal/http/handler/v1/health"
	"github.com/PestovOleg/mini-bank/backend/services/uproxy/internal/http/handler/v1/uproxy"
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	Handler     http.Handler
	Feature     string
	Middlewares []mux.MiddlewareFunc
}

func BaseRoutes(s *Services) map[string]map[string]RouteConfig {
	return map[string]map[string]RouteConfig{
		"/uproxy-minibank-health": {
			http.MethodGet: {
				Handler: health.NewHealthCheckHandler(),
			},
		},
		"/uproxy": {
			http.MethodGet: {
				Handler:     uproxy.NewUnleashProxyHandler().ListToggles(),
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware},
			},
		},
	}
}

func SetHandler(r *mux.Router, paths map[string]map[string]RouteConfig) {
	for path, methods := range paths {
		for method, config := range methods {
			handler := config.Handler
			for _, middleware := range config.Middlewares {
				handler = middleware(handler)
			}

			if config.Feature == "" {
				r.Handle(path, handler).Methods(method)
			} else {
				r.Handle(path, middleware.FeatureToggleMiddleware(config.Feature, handler)).Methods(method)
			}
		}
	}
}
