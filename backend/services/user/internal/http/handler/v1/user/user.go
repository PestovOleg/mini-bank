// TODO: сделать  описание
package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/PestovOleg/mini-bank/backend/services/user/domain/user"
	"github.com/PestovOleg/mini-bank/backend/services/user/internal/http/mapper"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type UserHandler struct {
	logger  *zap.Logger
	service *user.Service
}

func NewUserHandler(s *user.Service) *UserHandler {
	return &UserHandler{
		logger:  logger.GetLogger("UserAPI"),
		service: s,
	}
}

// UserCreateRequest represents the request payload for user creation.
// swagger:model
type UserCreateRequest struct {
	ID         string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" validate:"required,uuid"`
	Email      string `json:"email" example:"Ivanych@gmail.com" validate:"required,email"`
	Phone      string `json:"phone" example:"+7(495)999-99-99" validate:"required"`
	Birthday   string `json:"birthday" example:"02.01.2006" validate:"required"`
	Name       string `json:"name" example:"Ivan" validate:"required"`
	LastName   string `json:"last_name" example:"Ivanov" validate:"required"`
	Patronymic string `json:"patronymic" example:"Ivanych" validate:"required"`
}

// UserUpdateRequest represents the request payload for user update.
// swagger:model
type UserUpdateRequest struct {
	Email string `json:"email" example:"Ivanych@gmail.com" validate:"required,email"`
	Phone string `json:"phone" example:"+7(495)999-99-99"`
}

//nolint:gochecknoglobals
var validate *validator.Validate

// CreateUser godoc
// @Version 1.0
// @title CreateUser
// @Summary Create a new user
// @Description Create a new user using the provided details
// @tags user-minibank
// @Accept  json
// @Produce  json
// @Param user body UserCreateRequest true "User details for creation"
// @Success 201 {string} string "A new user has been created with ID: {id}"
// @Error 404 {string} "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
//
//nolint:gocognit
func (u *UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input UserCreateRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to decode request"))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		// валидация входных данных
		validate = validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok { //nolint:errorlint
				u.logger.Error("Validation error:" + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				_, err = w.Write([]byte("Unable to validate request:" + err.Error()))
				if err != nil {
					u.logger.Error(err.Error())
				}

				return
			}

			for _, err := range err.(validator.ValidationErrors) { //nolint:forcetypeassert,errorlint
				u.logger.Error("Validation Error, Field:" + err.Field() + " is " + err.ActualTag())
			}

			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to validate request:" + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		birthday, err := time.Parse("02.01.2006", input.Birthday)
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to parse birthday"))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		id, err := uuid.Parse(input.ID)
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		_, err = u.service.CreateUser(
			id,
			input.Email,
			input.Phone,
			input.Name,
			input.LastName,
			input.Patronymic,
			birthday,
		)

		u.logger.Debug(id.String())
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to create a new user: " + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		toJSON := &struct {
			ID string `json:"id"`
		}{
			ID: id.String(),
		}

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading ID"))
			if err != nil {
				u.logger.Error(err.Error())
			}
		}
		u.logger.Sugar().Infof("New user was created with ID: ", id.String())
	})
}

// GetUser godoc
// @title Get User by ID
// @version 1.0
// @summary Retrieve user details based on the provided ID.
// @description Fetch the user details using the provided user ID.
// @tags user-minibank
// @accept json
// @produce json
// @param id path string true "User ID"
// @success 200 {object} mapper.User "Successfully retrieved user details"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "User not found"
// @Security BasicAuth
// @router /users/{id} [get]
func (u *UserHandler) GetUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		data, err := u.service.GetUser(id)
		if err != nil && !errors.Is(err, user.ErrNotFound) {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		if errors.Is(err, user.ErrNotFound) || data == nil {
			u.logger.Sugar().Errorf("User with ID: %v not found", id)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		toJSON := &mapper.User{
			ID:         data.ID,
			Email:      data.Email,
			Phone:      data.Phone,
			Birthday:   data.Birthday.Format("02.01.2006"),
			Name:       data.Name,
			LastName:   data.LastName,
			Patronymic: data.Patronymic,
			CreatedAt:  data.CreatedAt.Format("02.01.2006"),
			UpdatedAt:  data.UpdatedAt.Format("02.01.2006"),
		}

		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading user"))
			if err != nil {
				u.logger.Error(err.Error())
			}
		}
	})
}

// UpdateUser godoc
// @title Update User by ID
// @version 1.0
// @summary Update user details based on the provided ID.
// @description Update the user details using the provided user ID.
// @tags user-minibank
// @accept json
// @produce json
// @param id path string true "User ID"
// @param user body UserUpdateRequest true "User Update Payload"
// @success 200 {string} string "Successfully updated user details"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "User not found"
// @Security BasicAuth
// @router /users/{id} [put]
func (u *UserHandler) UpdateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}
		var input UserUpdateRequest

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't decode request"))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		// валидация входных данных
		validate = validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok { //nolint:errorlint
				u.logger.Error("Validation error:" + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				_, err = w.Write([]byte("Unable to validate request:" + err.Error()))
				if err != nil {
					u.logger.Error(err.Error())
				}

				return
			}

			for _, err := range err.(validator.ValidationErrors) { //nolint:forcetypeassert,errorlint
				u.logger.Error("Validation Error, Field:" + err.Field() + " is " + err.ActualTag())
			}

			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to validate request:" + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		err = u.service.UpdateUser(id, input.Email, input.Phone)
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't update user: " + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}
		u.logger.Sugar().Infof("User %v was updated", id)
	})
}
