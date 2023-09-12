package health

import (
	"io"
	"net/http"

	"github.com/PestovOleg/mini-bank/pkg/util"
	"go.uber.org/zap"
)

type healthCheckHandler struct {
	logger *zap.Logger
}

func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{logger: util.GetLogger("API")}
}

func (h *healthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Service is healthy - Hello from Health Check Handler Endpoint")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
