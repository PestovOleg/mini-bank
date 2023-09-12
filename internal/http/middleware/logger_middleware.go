package middleware

import (
	"net/http"
	"time"

	"github.com/PestovOleg/mini-bank/pkg/util"
)

// Замена ответа для перехвата статуса обертываемого сервиса
type captureResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *captureResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := util.GetSugaredLogger("API")
		start := time.Now()
		captureWriter := &captureResponseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(captureWriter, r)
		statusCode := captureWriter.status
		logger.Infof(
			"Method: %s, Path: %s, Duration: %s, Status code: %d",
			r.Method,
			r.URL.Path,
			time.Since(start),
			statusCode,
		)
	})
}
