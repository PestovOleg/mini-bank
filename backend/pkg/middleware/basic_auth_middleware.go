package middleware

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
)

// TODO: покрыть логами
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLogger("BAuthMdle")
		client := &http.Client{
			Timeout: time.Second * 3,
		}
		host := os.Getenv("AUTH_VERIFY_HOST")

		if host == "" {
			logger.Error("sysvar AUTH_HOST is not enabled, URL to Auth server cannot be found")
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		authRequest, err := http.NewRequestWithContext(context.Background(), http.MethodPost, host, nil)
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		authRequest.Header.Set("Authorization", r.Header.Get("Authorization"))

		resp, err := client.Do(authRequest)
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusUnauthorized {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Error(err.Error())
				http.Error(w, "Internal server error", http.StatusInternalServerError)

				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			_, err = w.Write(body)
			if err != nil {
				logger.Error(err.Error())
			}

			return
		}

		next.ServeHTTP(w, r)
	})
}
