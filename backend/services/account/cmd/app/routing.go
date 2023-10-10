// Пакет роутов версии 1
// BaseRoutes - маршруты не содержащие middleware
//
//	[L] - включающие middleware: LoggingMiddleware
//	[LA]- включающие middleware: LoggingMiddleware,AuthorizeMiddleware
package app

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/middleware"
	handlerAccount "github.com/PestovOleg/mini-bank/backend/services/account/internal/http/handler/v1/account"
	"github.com/PestovOleg/mini-bank/backend/services/account/internal/http/handler/v1/health"
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	Handler     http.Handler
	Feature     string
	Middlewares []mux.MiddlewareFunc
}

func BaseRoutes(s *Services) map[string]map[string]RouteConfig {
	return map[string]map[string]RouteConfig{
		"/account-minibank-health": {
			http.MethodGet: {
				Handler: health.NewHealthCheckHandler(),
			},
		},
		"/users/{userid}/accounts": {
			http.MethodPost: {
				Handler:     handlerAccount.NewAccountHandler(s.AccountService).CreateAccount(),
				Feature:     "CreateAccountToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware, middleware.BasicAuthMiddleware},
			},
			http.MethodGet: {
				Handler:     handlerAccount.NewAccountHandler(s.AccountService).ListAccountsByUserID(),
				Feature:     "ListAccountsToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware, middleware.BasicAuthMiddleware},
			},
		},
		"/users/{userid}/accounts/{id}": {
			http.MethodPut: {
				Handler:     handlerAccount.NewAccountHandler(s.AccountService).UpdateAccount(),
				Feature:     "UpdateAccountToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware, middleware.BasicAuthMiddleware},
			},
			http.MethodGet: {
				Handler:     handlerAccount.NewAccountHandler(s.AccountService).GetAccountByID(),
				Feature:     "GetAccountToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware, middleware.BasicAuthMiddleware},
			},
			http.MethodDelete: {
				Handler:     handlerAccount.NewAccountHandler(s.AccountService).DeleteAccount(),
				Feature:     "DeleteAccountToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware, middleware.BasicAuthMiddleware},
			},
		},
		"/users/{userid}/accounts/{id}/topup": {
			http.MethodPut: {
				Handler:     handlerAccount.NewAccountHandler(s.AccountService).TopUp(),
				Feature:     "TopUpToggle",
				Middlewares: []mux.MiddlewareFunc{middleware.LoggerMiddleware, middleware.BasicAuthMiddleware},
			},
		},
		"/users/{userid}/accounts/{id}/withdraw": {
			http.MethodPut: {
				Handler:     handlerAccount.NewAccountHandler(s.AccountService).Withdraw(),
				Feature:     "WithdrawToggle",
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
