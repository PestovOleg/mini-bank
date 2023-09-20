package middleware

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/domain/user"
	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/gorilla/mux"
)

func BasicAuthMiddleware(s *user.Service) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username, pass, ok := r.BasicAuth()
			if !ok {
				unauthorized(w)

				return
			}

			u, err := s.GetUserByUName(username)
			if err != nil {
				unauthorized(w)

				return
			}

			err = u.VerifyPassword(pass)
			if err != nil {
				unauthorized(w)

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func unauthorized(w http.ResponseWriter) {
	realm := "Предоставьте данные для авторизации"
	w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
	w.WriteHeader(http.StatusUnauthorized)
	_, err := w.Write([]byte("Unauthorized.\n"))

	if err != nil {
		logger := logger.GetLogger("API")
		logger.Error(err.Error())
	}
}
