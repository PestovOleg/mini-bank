package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"../../backend/services/user/domain/user"
	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/PestovOleg/mini-bank/backend/services/user/internal/http/mapper"
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
		logger:  logger.GetLogger("API"),
		service: s,
	}
}

// UserCreateRequest represents the request payload for user creation.
// swagger:model
type UserCreateRequest struct {
	ID         string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email      string `json:"email" example:"Ivanych@gmail.com"`
	Phone      string `json:"phone" example:"+7(495)999-99-99"`
	Birthday   string `json:"birthday" example:"2013-Feb-03"`
	Name       string `json:"name" example:"Ivan"`
	LastName   string `json:"lastName" example:"Ivanov"`
	Patronymic string `json:"patronymic" example:"Ivanych"`
}

// UserUpdateRequest represents the request payload for user update.
// swagger:model
type UserUpdateRequest struct {
	Email string `json:"email" example:"Ivanych@gmail.com"`
	Phone string `json:"phone" example:"+7(495)999-99-99"`
}

// CreateUser godoc
// @Version 1.0
// @title CreateUser
// @Summary Create a new user
// @Description Create a new user using the provided details
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body UserCreateRequest true "User details for creation"
// @Success 201 {string} string "A new user has been created with ID: {id}"
// @Error 404 {string} "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
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
		birthday, err := time.Parse(time.RFC3339, input.Birthday)
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

		err = u.service.CreateUser(
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

		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading ID"))
			if err != nil {
				u.logger.Error(err.Error())
			}
		}
		u.logger.Sugar().Infof("New user was created with ID: ", id.String())
		w.WriteHeader(http.StatusCreated)
	})
}

// GetUser godoc
// @title Get User by ID
// @version 1.0
// @summary Retrieve user details based on the provided ID.
// @description Fetch the user details using the provided user ID.
// @tags users
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
			Username:   data.Username,
			Email:      data.Email,
			Phone:      data.Phone,
			Birthday:   data.Birthday.Format(time.RFC3339),
			Name:       data.Name,
			LastName:   data.LastName,
			Patronymic: data.Patronymic,
			IsActive:   data.IsActive,
			CreatedAt:  data.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  data.UpdatedAt.Format(time.RFC3339),
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
// @tags users
// @accept json
// @produce json
// @param id path string true "User ID"
// @param body UserUpdateRequest true "User Update Payload"
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
		var input struct {
			Email string `json:"email"`
			Phone string `json:"phone"`
		}

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

// DeleteUser godoc
// @title Delete User by ID
// @version 1.0
// @summary Delete user based on the provided ID.
// @description Delete the user using the provided user ID.
// @tags users
// @accept json
// @produce json
// @param id path string true "User ID"
// @success 200 {string} string "Successfully deleted user"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "User not found"
// @Security BasicAuth
// @router /users/{id} [delete]
func (u *UserHandler) DeleteUser() http.Handler {
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

		err = u.service.DeleteUser(id)
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't delete user: " + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}
		u.logger.Sugar().Infof("User %v was deleted", id)
	})
}

// Enter godoc
// @title Enter with credentials
// @version 1.0
// @summary Get User ID with credentials.
// @description Get User ID with credentials.
// @tags users
// @accept json
// @produce json
// @success 200 {string} ID "Successfully retrieved User ID"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "User not found"
// @Security BasicAuth
// @router /users [get]
func (u *UserHandler) Enter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, _, _ := r.BasicAuth()
		user, _ := u.service.GetUserByUName(username)

		w.Header().Set("Content-Type", "application/json")
		toJSON := &struct {
			ID string `json:"id"`
		}{
			ID: user.ID.String(),
		}
		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading ID"))
			if err != nil {
				u.logger.Error(err.Error())
			}
		}
	})
}
