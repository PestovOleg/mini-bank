// Пакет роутов версии 1
// BaseRoutes - маршруты не содержащие middleware
//
//	[L] - включающие middleware: LoggingMiddleware
//	[LA]- включающие middleware: LoggingMiddleware,AuthorizeMiddleware
package app

import (
	"net/http"

	handlerAccount "github.com/PestovOleg/mini-bank/backend/internal/http/handler/v1/account"
	"github.com/PestovOleg/mini-bank/backend/internal/http/handler/v1/health"
	handlerUser "github.com/PestovOleg/mini-bank/backend/internal/http/handler/v1/user"
	"github.com/PestovOleg/mini-bank/backend/internal/http/middleware"
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	Handler http.Handler
	Feature string
}

func BaseRoutes(s *Services) map[string]map[string]RouteConfig {
	return map[string]map[string]RouteConfig{
		"/health": {
			http.MethodGet: {
				Handler: health.NewHealthCheckHandler(),
			},
		},
		"/users": {
			http.MethodPost: {
				Handler: handlerUser.NewUserHandler(s.UserService).CreateUser(),
				Feature: "CreateUserToggle",
			},
		},
	}
}

func BaseRoutesL(s *Services) map[string]map[string]RouteConfig {
	return map[string]map[string]RouteConfig{
		"/users": {
			http.MethodGet: {
				Handler: handlerUser.NewUserHandler(s.UserService).Enter(),
			},
		},
		"/users/{id}": {
			http.MethodGet: {
				Handler: handlerUser.NewUserHandler(s.UserService).GetUser(),
				Feature: "GetUserToggle",
			},
			http.MethodPut: {
				Handler: handlerUser.NewUserHandler(s.UserService).UpdateUser(),
				Feature: "UpdateUserToggle",
			},
			http.MethodDelete: {
				Handler: handlerUser.NewUserHandler(s.UserService).DeleteUser(),
				Feature: "DeleteUserToggle",
			},
		},
		"/users/{id}/accounts": {
			http.MethodPost: {
				Handler: handlerAccount.NewAccountHandler(s.AccountService).CreateAccount(),
				Feature: "CreateAccountToggle",
			},
			http.MethodGet: {
				Handler: handlerAccount.NewAccountHandler(s.AccountService).ListAccountsByUserID(),
				Feature: "ListAccountsToggle",
			},
		},
		"/users/{userid}/accounts/{id}": {
			http.MethodPut: {
				Handler: handlerAccount.NewAccountHandler(s.AccountService).UpdateAccount(),
				Feature: "UpdateAccountToggle",
			},
			http.MethodGet: {
				Handler: handlerAccount.NewAccountHandler(s.AccountService).GetAccountByID(),
				Feature: "GetAccountToggle",
			},
			http.MethodDelete: {
				Handler: handlerAccount.NewAccountHandler(s.AccountService).DeleteAccount(),
				Feature: "DeleteAccountToggle",
			},
		},
		"/users/{userid}/accounts/{id}/topup": {
			http.MethodPut: {
				Handler: handlerAccount.NewAccountHandler(s.AccountService).TopUp(),
				Feature: "TopUpToggle",
			},
		},
		"/users/{userid}/accounts/{id}/withdraw": {
			http.MethodPut: {
				Handler: handlerAccount.NewAccountHandler(s.AccountService).Withdraw(),
				Feature: "WithdrawToggle",
			},
		},
	}
}

func SetHandler(r *mux.Router, paths map[string]map[string]RouteConfig) {
	for path, methods := range paths {
		for method, config := range methods {
			if config.Feature == "" {
				r.Handle(path, config.Handler).Methods(method)
			} else {
				r.Handle(path, middleware.FeatureToggleMiddleware(config.Feature, config.Handler)).Methods(method)
			}
		}
	}
}