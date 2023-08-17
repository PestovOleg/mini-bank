package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/PestovOleg/mini-bank/pkg/util"
	"go.uber.org/zap"
)

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
