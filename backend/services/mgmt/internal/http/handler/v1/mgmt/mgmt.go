// TODO: сделать  описание
package mgmt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	Username   string `json:"username" example:"Ivanec" validate:"required,latinusername"`
	Password   string `json:"password" example:"mypass" validate:"required"`
	Email      string `json:"email" example:"Ivanych@gmail.com" validate:"required,email"`
	Phone      string `json:"phone" example:"+7(495)999-99-99" validate:"required"`
	Birthday   string `json:"birthday" example:"02.01.2006" validate:"required"`
	Name       string `json:"name" example:"Ivan" validate:"required"`
	LastName   string `json:"last_name" example:"Ivanov" validate:"required"`
	Patronymic string `json:"patronymic" example:"Ivanych" validate:"required"`
}

//nolint:gochecknoglobals
var validate *validator.Validate

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

		// валидация входных данных
		validate = validator.New()

		validate.RegisterValidation("latinusername", func(fl validator.FieldLevel) bool { //nolint:errcheck
			// только латинские символы и цифры
			return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(fl.Field().String())
		})

		err = validate.Struct(input)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok { //nolint:errorlint
				m.logger.Error("Validation error:" + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				_, err = w.Write([]byte("Unable to validate request:" + err.Error()))
				if err != nil {
					m.logger.Error(err.Error())
				}

				return
			}

			for _, err := range err.(validator.ValidationErrors) { //nolint:forcetypeassert,errorlint
				m.logger.Error("Validation Error, Field:" + err.Field() + " is " + err.ActualTag())
			}

			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to validate request:" + err.Error()))
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
			http.Error(w, "Internal server error", http.StatusInternalServerError) // TODO: сделать возврат ошибок

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

// DeleteUser godoc
// @Version 1.0
// @title CreateUser
// @Summary Orchestrate deactivation of a  user with services auth, user and account
// @Description Delete(deactivate) a  user using the provided details
// @Tags mgmt
// @Accept  json
// @Produce  json
// @param id path string true "User ID"
// @Success 200 {string} string "The user has been deactivated"
// @Error 404 {string} "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Router /mgmt/{id} [delete]
// @Security BasicAuth
//
//nolint:gocognit
func (m *MgmtHandler) DeleteUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, err := uuid.Parse(vars["id"])
		if err != nil {
			m.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				m.logger.Error(err.Error())
			}

			return
		}

		// Часть 1. Получаем список счетов пользователя
		client := &http.Client{
			Timeout: time.Second * 3,
		}
		authHost := os.Getenv("AUTH_HOST")
		accountHost := os.Getenv("ACCOUNT_HOST")
		verifyHost := os.Getenv("AUTH_VERIFY_HOST")

		if authHost == "" || accountHost == "" || verifyHost == "" {
			m.logger.Error(
				"sysvar AUTH_HOST or ACCOUNT_HOST or AUTH_VERIFY_HOST," +
					"is not enabled,URL to Auth and Account services cannot be found",
			)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}
		m.logger.Error("Sysvars 're OK")
		accountListRequest, err := http.NewRequestWithContext(
			r.Context(),
			http.MethodGet,
			accountHost+"/"+userID.String()+"/accounts", nil,
		)
		if err != nil {
			m.logger.Sugar().Error("Failed to create new HTTP request:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		accountListRequest.Header.Set("Authorization", r.Header.Get("Authorization"))
		accountListRequest.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(accountListRequest)
		if err != nil {
			m.logger.Sugar().Error("Failed to send HTTP request:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		// Если статус отличен от успешного - возвращаем ошибку
		if resp.StatusCode != http.StatusOK {
			m.logger.Debug("Status code of response (account) = " + strconv.Itoa(resp.StatusCode))
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				m.logger.Sugar().Error("Failed to get account list: %s", err.Error())
				http.Error(w, "Internal server error", http.StatusInternalServerError)

				return
			}

			m.logger.Sugar().Error("Failed to get account list, body: %s", string(body))
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		m.logger.Debug("Список счетов получен")

		// Чтение ответа
		type AccountListResponse struct {
			ID           string    `json:"id"`
			UserID       string    `json:"user_id"`
			Account      string    `json:"account"`
			Currency     string    `json:"currency"`
			Name         string    `json:"name"`
			Amount       float64   `json:"amount"`
			InterestRate float64   `json:"interest_rate"`
			IsActive     bool      `json:"is_active"`
			CreatedAt    time.Time `json:"created_at"`
			UpdatedAt    time.Time `json:"updated_at"`
		}

		var accounts []AccountListResponse

		body, _ := io.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &accounts); err != nil {
			m.logger.Error("Failed to get unmarshall account list, body:")
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}
		resp.Body.Close()

		// Часть 2. Закрываем все счета клиента

		var host string
		for _, account := range accounts {
			host = accountHost + "/" + account.UserID + "/accounts/" + account.ID
			accountDeleteRequest, err := http.NewRequestWithContext(
				r.Context(),
				http.MethodDelete,
				host, nil,
			)
			if err != nil {
				m.logger.Sugar().Error("Failed to create new HTTP request for deleting account:", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)

				return
			}

			accountDeleteRequest.Header.Set("Authorization", r.Header.Get("Authorization"))
			accountDeleteRequest.Header.Set("Content-Type", "application/json")

			// делаем запрос в account сервис
			resp, err := client.Do(accountDeleteRequest)
			if err != nil {
				m.logger.Sugar().Error("Failed to send HTTP request to delete account:", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)

				return
			}

			// Если статус отличен от успешного - возвращаем ошибку
			if resp.StatusCode != http.StatusOK {
				m.logger.Debug("Status code of response (account) = " + strconv.Itoa(resp.StatusCode))
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					m.logger.Sugar().Error("Failed to delete account: %s", err.Error())
					http.Error(w, "Internal server error", http.StatusInternalServerError)

					return
				}

				m.logger.Sugar().Error("Failed to delete account, body: %s", string(body))
				http.Error(w, "Internal server error", http.StatusInternalServerError)

				return
			}
			m.logger.Debug("Счет удален: " + account.ID)
			resp.Body.Close()
		}

		// Часть 3,удаляем самого клиента
		userDeleteRequest, err := http.NewRequestWithContext(
			r.Context(),
			http.MethodDelete,
			authHost+"/"+userID.String(), nil,
		)
		if err != nil {
			m.logger.Sugar().Error("Failed to create new HTTP request:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		userDeleteRequest.Header.Set("Authorization", r.Header.Get("Authorization"))
		userDeleteRequest.Header.Set("Content-Type", "application/json")

		// делаем запрос в auth сервис
		resp, err = client.Do(userDeleteRequest)
		if err != nil {
			m.logger.Sugar().Error("Failed to send HTTP request:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		// Если статус отличен от успешного - возвращаем ошибку
		if resp.StatusCode != http.StatusOK {
			m.logger.Debug("Status code of response (auth) = " + strconv.Itoa(resp.StatusCode))
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				m.logger.Sugar().Error("Failed to delete user: %s", err.Error())
				http.Error(w, "Internal server error", http.StatusInternalServerError)

				return
			}

			m.logger.Sugar().Error("Failed to delete user, body: %s", string(body))
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}
		resp.Body.Close()
		m.logger.Debug("User was deleted: " + userID.String())
	})
}
