package middleware

import (
	"context"
	"io"
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
)

// TODO: покрыть логами
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLogger("BAuthMdle")
		client := &http.Client{}
		authRequest, err := http.NewRequestWithContext(context.Background(), "GET", "http://localhost/api/v1/auth", nil)
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
			// If unauthorized, copy the response from the authorization request to the original response writer
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Error(err.Error())
				http.Error(w, "Internal server error", http.StatusInternalServerError)

				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(body)
			return
		}

		next.ServeHTTP(w, r)
	})
}
