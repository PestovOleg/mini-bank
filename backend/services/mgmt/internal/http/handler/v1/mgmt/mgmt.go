// TODO: сделать  описание
package mgmt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"go.uber.org/zap"
)

type MgmtHandler struct {
	logger *zap.Logger
}

func NewMgmtHandler() *MgmtHandler {
	return &MgmtHandler{
		logger: logger.GetLogger("MgmtAPI"),
	}
}

// UserCreateRequest represents the request payload for user creation.
// swagger:model
type MgmtCreateUserRequest struct {
	Username   string `json:"username" example:"Ivanec"`
	Password   string `json:"password" example:"mypass"`
	Email      string `json:"email" example:"Ivanych@gmail.com"`
	Phone      string `json:"phone" example:"+7(495)999-99-99"`
	Birthday   string `json:"birthday" example:"02.01.2006"`
	Name       string `json:"name" example:"Ivan"`
	LastName   string `json:"last_name" example:"Ivanov"`
	Patronymic string `json:"patronymic" example:"Ivanych"`
}

// CreateUser godoc
// @Version 1.0
// @title CreateUser
// @Summary Orchestrate creation of a new user with services auth and user
// @Description Create a new user using the provided details
// @Tags mgmt
// @Accept  json
// @Produce  json
// @Param user body MgmtCreateUserRequest true "User details for creation"
// @Success 201 {string} string "A new user has been created with ID: {id}"
// @Error 404 {string} "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Router /mgmt [post]
//
//nolint:gocognit
func (m *MgmtHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input MgmtCreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			m.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to decode request"))
			if err != nil {
				m.logger.Error(err.Error())
			}

			return
		}

		// Часть 1. Вызов auth
		// готовим запрос в сервис auth
		client := &http.Client{
			Timeout: time.Second * 5,
		}
		authHost := os.Getenv("AUTH_HOST")
		userHost := os.Getenv("USER_HOST")

		if authHost == "" || userHost == "" {
			m.logger.Error("sysvar AUTH_HOST or USER_HOST is not enabled, URL to Auth and User services cannot be found")
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		type AuthRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		authData := AuthRequest{
			Username: input.Username,
			Password: input.Password,
		}

		jsonData, err := json.Marshal(authData)
		if err != nil {
			m.logger.Sugar().Error("Failed to marshal JSON:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		authRequest, err := http.NewRequestWithContext(
			r.Context(),
			http.MethodPost,
			authHost,
			bytes.NewBuffer(jsonData),
		)
		if err != nil {
			m.logger.Sugar().Error("Failed to create new HTTP request:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		authRequest.Header.Set("Content-Type", "application/json")

		// делаем запрос в auth сервис
		resp, err := client.Do(authRequest)
		if err != nil {
			m.logger.Sugar().Error("Failed to send HTTP request:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		// Если статус отличен от успешного - возвращаем ошибку
		if resp.StatusCode != http.StatusCreated {
			m.logger.Debug("Status code of response (auth) = " + strconv.Itoa(resp.StatusCode))
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				m.logger.Sugar().Error("Failed to create user: %s", err.Error())
				http.Error(w, "Internal server error", http.StatusInternalServerError)

				return
			}

			m.logger.Sugar().Error("Failed to create user, body: %s", string(body))
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		// Чтение ответа
		type AuthResponse struct {
			ID string `json:"id"`
		}

		var authResp AuthResponse

		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&authResp); err != nil {
			m.logger.Sugar().Error("Failed to decode JSON response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError) // TODO:сделать возврат ошибок

			return
		}
		resp.Body.Close()
		// Часть 2. Вызов сервиса User

		type UserCreateRequest struct {
			ID         string `json:"id"`
			Email      string `json:"email"`
			Phone      string `json:"phone"`
			Birthday   string `json:"birthday"`
			Name       string `json:"name"`
			LastName   string `json:"last_name"`
			Patronymic string `json:"patronymic"`
		}

		userData := UserCreateRequest{
			ID:         authResp.ID,
			Email:      input.Email,
			Phone:      input.Phone,
			Birthday:   input.Birthday,
			Name:       input.Name,
			LastName:   input.LastName,
			Patronymic: input.Patronymic,
		}

		jsonData, err = json.Marshal(userData)
		if err != nil {
			m.logger.Sugar().Error("Failed to marshal user JSON:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}
		userRequest, err := http.NewRequestWithContext(
			r.Context(),
			http.MethodPost,
			userHost,
			bytes.NewBuffer(jsonData),
		)
		if err != nil {
			m.logger.Sugar().Error("Failed to create user HTTP request:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		userRequest.Header.Set("Content-Type", "application/json")
		resp, err = client.Do(userRequest)
		if err != nil {
			m.logger.Sugar().Error("Failed to send user HTTP request:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		// Если статус отличен от успешного - возвращаем ошибку
		// TODO: сделать удаление в auth
		if resp.StatusCode != http.StatusCreated {
			m.logger.Debug("Status code of response (user) = " + strconv.Itoa(resp.StatusCode))
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				m.logger.Sugar().Error("Failed to create user: %s", err.Error())
				http.Error(w, "Internal server error", http.StatusInternalServerError)

				return
			}

			m.logger.Sugar().Error("Failed to create user, body: %s", string(body))
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		var userResp struct {
			ID string `json:"id"`
		}

		dec = json.NewDecoder(resp.Body)
		if err := dec.Decode(&userResp); err != nil {
			m.logger.Sugar().Error("Failed to decode user JSON response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}
		resp.Body.Close()

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(userResp); err != nil {
			m.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading ID"))
			if err != nil {
				m.logger.Error(err.Error())
			}
		}
		m.logger.Sugar().Infof("New user was created with ID: ", userResp.ID)
	})
}
