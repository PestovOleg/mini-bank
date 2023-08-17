package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func BaseRoutes(l *zap.Logger) map[string]map[string]http.Handler {
	return map[string]map[string]http.Handler{
		"/health-check": {http.MethodGet: NewHealthCheckHandler()},
	}
}

func SetHandler(r *mux.Router, paths map[string]map[string]http.Handler) {
	for path := range paths {
		for method, handler := range paths[path] {
			r.Handle(path, handler).Methods(method)
		}
	}
}
