package health

import (
	"io"
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"go.uber.org/zap"
)

type healthCheckHandler struct {
	logger *zap.Logger
}

func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{logger: logger.GetLogger("API")}
}

// HealthCheck godoc
// @title		 Health Check
// @version 	 1.0
// @summary      Check the health status of the server
// @description  Returns the server's health status.
// @tags         account-minibank
// @success 200 {string} string "Service is healthy - Hello from Health Check Handler Endpoint" "StatusOK"
// @example 200 {string} "Service is healthy - Hello from Health Check Handler Endpoint"
// @failure 500 {string} string "StatusInternalError"
// @router       /account-minibank-health [get]
func (h *healthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Account Service is healthy - Hello from Health Check Handler Endpoint")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
