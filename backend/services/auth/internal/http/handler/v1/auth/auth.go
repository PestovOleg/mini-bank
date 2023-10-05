package auth

import (
	"encoding/json"
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/PestovOleg/mini-bank/backend/services/auth/domain/auth"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type AuthHandler struct {
	logger  *zap.Logger
	service *auth.Service
}

func NewAuthHandler(s *auth.Service) *AuthHandler {
	return &AuthHandler{
		logger:  logger.GetLogger("AuthAPI"),
		service: s,
	}
}

// AuthCreateRequest represents the request payload for authentication record creation.
// swagger:model
type AuthCreateRequestquest struct {
	Username string `json:"username" example:"Ivanec"`
	Password string `json:"password" example:"mypass"`
}

// AuthWithCredentials represents the request payload for authentication with credentials.
// swagger:model
type AuthWithCredentials struct {
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
}

// AuthWithToken represents the request payload for Basic authorization .
// swagger:model
type AuthWithToken struct {
	Token string `json:"token" example:"Basic dXNlcjE6cGFzc3dvcmQx"`
}

// CreateAuth godoc
// @Version 1.0
// @title CreateAuth
// @Summary Create a new authentication record
// @Description Create a new authentication record using the provided details
// @Tags authentication
// @Accept  json
// @Produce  json
// @Param user body AuthCreateRequest true "Authentication details for creation"
// @Success 201 {string} string "A new authentication record has been created with ID: {id}"
// @Error 404 {string} "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Router /auth [post]
func (a *AuthHandler) CreateAuth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input AuthCreateRequestquest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to decode request"))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		id, err := a.service.CreateAuth(
			input.Username,
			input.Password,
		)

		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to create a new user: " + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		a.logger.Debug("Record created with ID: " + id.String())

		toJSON := &struct {
			ID string `json:"id"`
		}{
			ID: id.String(),
		}

		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading ID"))
			if err != nil {
				a.logger.Error(err.Error())
			}
		}
		a.logger.Sugar().Infof("New authentication record was created with ID: ", id.String())
		w.WriteHeader(http.StatusCreated)
	})
}

// DeleteAuth godoc
// @title Deactivate Authentication record by ID
// @version 1.0
// @summary Deactivate authentication record based on the provided ID.
// @description Deactivate the authentication record using the provided user ID.
// @tags auth
// @accept json
// @produce json
// @param id path string true "Auth ID"
// @success 200 {string} string "Successfully deleted authentication record"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "Page not found"
// @Security BasicAuth
// @router /auth/{id} [delete]
func (a *AuthHandler) DeleteAuth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		err = a.service.DeleteAuth(id)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't delete user: " + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}
		a.logger.Sugar().Infof("Authentication record %v was deleted", id)
	})
}

// Authenticate godoc
// @title Authenticate with credentials
// @version 1.0
// @summary Authenticate User with credentials.
// @description Get User ID with credentials.
// @tags users
// @accept json
// @produce json
// @success 200 {string} ID "Successfully retrieved User ID"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "User not found"
// @Security BasicAuth
// @router /users [get]
func (a *AuthHandler) Authenticate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, pass, _ := r.BasicAuth()
		id, err := a.service.AuthenticateUser(username, pass)
		if err != nil {
			unauthorized(w)
		}

		w.Header().Set("Content-Type", "application/json")
		toJSON := &struct {
			ID string `json:"id"`
		}{
			ID: id.String(),
		}
		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading ID"))
			if err != nil {
				a.logger.Error(err.Error())
			}
		}
	})
}

// Authorize godoc
// @title Authorize with token
// @version 1.0
// @summary Authorize User with token.
// @description Authorize User with token.
// @tags users
// @accept json
// @produce json
// @success 200 {string} ID "Authorized"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "User not found"
// @Security BasicAuth
// @router /auth [get]
func (a *AuthHandler) Authorize() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		err := a.service.AuthorizeUser(authHeader)
		if err != nil {
			unauthorized(w)
		}
	})
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
