package account

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/domain/account"
	"github.com/PestovOleg/mini-bank/backend/internal/http/mapper"
	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type AccountHandler struct {
	logger  *zap.Logger
	service *account.Service
}

func NewAccountHandler(s *account.Service) *AccountHandler {
	return &AccountHandler{
		logger:  logger.GetLogger("API"),
		service: s,
	}
}

// AccountCreateRequest represents the request payload for account creation.
// swagger:model
type AccountCreateRequest struct {
	Currency string `json:"currency" example:"810"`
}

// CreateAccount godoc
// @Version 1.0
// @title CreateAccount
// @Summary Create a new account
// @Description Create a new account using the provided details
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param user body AccountCreateRequest true "Account details for creation"
// @Success 201 {uuid} string "A new account has been created with number: {account}"
// @Error 404 {string} "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users/{id}/accounts [post]
func (u *AccountHandler) CreateAccount() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		userID, err := uuid.Parse(vars["id"])
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}
		var input AccountCreateRequest
		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to decode request"))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		account, err := u.service.CreateAccount(
			userID,
			input.Currency,
		)

		u.logger.Debug(account.Account)
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to create a new account: " + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}
		u.logger.Sugar().Infof("New account was created with number: %s", account.Account)
		w.WriteHeader(http.StatusCreated)
	})
}

// GetAccount godoc
// @title Get Account by ID
// @version 1.0
// @summary Retrieve account details based on the provided ID.
// @description Fetch the account details using the provided account ID.
// @tags accounts
// @accept json
// @produce json
// @param id path string true "Account ID"
// @success 200 {object} mapper.Account "Successfully retrieved account details"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "Account not found"
// @Security BasicAuth
// @router /users/{userid}/accounts/{id} [get]
func (u *AccountHandler) GetAccountByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		userID, err := uuid.Parse(vars["userid"])
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}
		// TODO: проверка этого ли клиента счет
		accountID, err := uuid.Parse(vars["id"])
		if err != nil {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		data, err := u.service.GetAccountByID(accountID)
		if err != nil && !errors.Is(err, account.ErrNotFound) {
			u.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		if errors.Is(err, account.ErrNotFound) || data == nil {
			u.logger.Sugar().Errorf("Account with ID: %v not found", accountID)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				u.logger.Error(err.Error())
			}

			return
		}

		toJSON := &mapper.Account{
			ID:           data.ID,
			UserID:       data.UserID,
			Account:      data.Currency,
			Currency:     data.Currency,
			Amount:       data.Amount,
			InterestRate: data.InterestRate,
			IsActive:     data.IsActive,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
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
