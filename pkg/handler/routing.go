package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/PestovOleg/mini-bank/pkg/util"
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

type healthCheckHandler struct {
	logger *zap.Logger
}

func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{logger: util.Getlogger("API")}
}

func (h *healthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_, err := io.WriteString(w, "Service is healthy - Hello from Health Check Handler Endpoint")

	h.logger.Info("Health check from " + r.RemoteAddr)

	if err != nil {
		fmt.Println("Some error in Health Check Handler Endpoint")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
