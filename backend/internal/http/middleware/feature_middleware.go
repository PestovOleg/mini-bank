package middleware

import (
	"net/http"

	"github.com/Unleash/unleash-client-go/v3"
)

func FeatureToggleMiddleware(featureName string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !unleash.IsEnabled(featureName) && featureName != "" {
			http.Error(w, "Feature not enabled", http.StatusForbidden)

			return
		}
		next.ServeHTTP(w, r)
	})
}
