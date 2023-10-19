package auth

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/PestovOleg/mini-bank/backend/services/auth/domain/auth"
	"github.com/go-playground/validator/v10"
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
type AuthCreateRequest struct {
	Username string `json:"username" example:"Ivanec" validate:"required,latinusername"`
	Password string `json:"password" example:"mypass" validate:"required"`
}

// AuthWithCredentials represents the request payload for authentication with credentials.
// swagger:model
type AuthWithCredentials struct {
	Username string `json:"username" example:"username" validate:"required,latinusername"`
	Password string `json:"password" example:"password" validate:"required"`
}

// AuthWithToken represents the request payload for Basic authorization .
// swagger:model
type AuthWithToken struct {
	Token string `json:"token" example:"Basic dXNlcjE6cGFzc3dvcmQx" validate:"required"`
}

//nolint:gochecknoglobals
var validate *validator.Validate

// CreateAuth godoc
// @Version 1.0
// @title CreateAuth
// @Summary Create a new authentication record (inter-service interaction)
// @Description Create a new authentication record using the provided details.
// @Tags auth-minibank
// @Accept  json
// @Produce  json
// @Param user body AuthCreateRequest true "Authentication details for creation"
// @Success 201 {string} string "A new authentication record has been created with ID: {id}"
// @Error 404 {string} "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Router /auth [post]
func (a *AuthHandler) CreateAuth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input AuthCreateRequest
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

		// валидация входных данных
		validate = validator.New()

		validate.RegisterValidation("latinusername", func(fl validator.FieldLevel) bool { //nolint:errcheck
			// только латинские символы и цифры
			return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(fl.Field().String())
		})

		err = validate.Struct(input)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok { //nolint:errorlint
				a.logger.Error("Validation error:" + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				_, err = w.Write([]byte("Unable to validate request:" + err.Error()))
				if err != nil {
					a.logger.Error(err.Error())
				}

				return
			}

			for _, err := range err.(validator.ValidationErrors) { //nolint:forcetypeassert,errorlint
				a.logger.Error("Validation Error, Field:" + err.Field() + " is " + err.ActualTag())
			}

			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to validate request:" + err.Error()))
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

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading ID"))
			if err != nil {
				a.logger.Error(err.Error())
			}
		}
		a.logger.Sugar().Infof("New authentication record was created with ID: ", id.String())
	})
}

// DeactivateAuth godoc
// @title Deactivate Authentication record by ID
// @version 1.0
// @summary Deactivate authentication record based on the provided ID.
// @description Deactivate the authentication record using the provided user ID.
// @tags auth-minibank
// @accept json
// @produce json
// @param id path string true "Auth ID"
// @success 200 {string} string "Successfully deleted authentication record"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "Page not found"
// @Security BasicAuth
// @router /auth/{id} [put]
func (a *AuthHandler) DeactivateAuth() http.Handler {
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

		err = a.service.DeactivateAuth(id)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't deactivate user: " + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}
		a.logger.Sugar().Infof("Authentication record %v was deactivated", id)
	})
}

// Authenticate godoc
// @title Authenticate with credentials
// @version 1.0
// @summary Authenticate User with credentials.
// @description Get User ID with credentials.
// @tags auth-minibank
// @accept json
// @produce json
// @success 200 {string} ID "Successfully retrieved User ID"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "User not found"
// @Security BasicAuth
// @router /auth/login [post]
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
// @tags auth-minibank
// @accept json
// @produce json
// @success 200 {string} ID "Authorized"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "User not found"
// @Security BasicAuth
// @router /auth/verify [post]
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

// DeleteAuth Delete Auth record, only inter-service communication
// @title Delete Authentication record by ID
// @version 1.0
// @summary Delete authentication record based on the provided ID.
// @description Delete the authentication record using the provided user ID.
// @tags auth-minibank
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
		a.logger.Debug("DeleteAuth was initiated")
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
			_, err = w.Write([]byte("Couldn't delete auth record: " + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}
		a.logger.Info("Authentication record was deleted")
	})
}
