package account

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/PestovOleg/mini-bank/backend/services/account/domain/account"
	"github.com/PestovOleg/mini-bank/backend/services/account/internal/http/mapper"
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
		logger:  logger.GetLogger("AccountAPI"),
		service: s,
	}
}

// AccountCreateRequest represents the request payload for account creation.
// swagger:model
type AccountCreateRequest struct {
	UserID   string `json:"user_id" example:"fdee7aae-f79f-4653-8a16-9207e6805b93"`
	Currency string `json:"currency" example:"810"`
	Name     string `json:"name" example:"Удачный"`
}

// AccountCreateRequest represents the request payload for account creation.
// swagger:model
type AccountUpdateRequest struct {
	Name         string  `json:"name" example:"Удачный"`
	InterestRate float64 `json:"interest_rate" example:"0.1250"`
}

// ChangeBalanceRequest represents the request payload for account creation.
// swagger:model
type ChangeBalanceRequest struct {
	Amount float64 `json:"amount" example:"9999.99"`
}

// CreateAccount godoc
// @Version 1.0
// @title CreateAccount
// @Summary Create a new account
// @Description Create a new account using the provided details
// @Tags account-minibank
// @Accept  json
// @Produce  json
// @Param user body AccountCreateRequest true "Account details for creation"
// @Success 201 {string} string "A new account has been created with number: {string}"
// @Error 404 {string} "Page not found"
// @Failure 500 {string} string "Internal server error"
// @Security BasicAuth
// @Router /accounts [post]
func (a *AccountHandler) CreateAccount() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var input AccountCreateRequest
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

		userID, err := uuid.Parse(input.UserID)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		a.logger.Sugar().Debugf("userID: %v, input.Currency: %s", userID, input.Currency)

		account, err := a.service.CreateAccount(
			userID,
			input.Currency,
			input.Name,
		)

		a.logger.Debug(account.Account)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Unable to create a new account: " + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		toJSON := &struct {
			ID string `json:"id"`
		}{
			ID: account.ID.String(),
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
		a.logger.Sugar().Infof("New account was created with number: %s", account.Account)
	})
}

// GetAccount godoc
// @title Get Account by ID
// @version 1.0
// @summary Retrieve account details based on the provided ID.
// @description Fetch the account details using the provided account ID.
// @tags account-minibank
// @accept json
// @produce json
// @param id path string true "Account ID"
// @param userid path string true "User ID"
// @success 200 {object} mapper.Account "Successfully retrieved account details"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "Account not found"
// @Security BasicAuth
// @router /accounts/{id} [get]
func (a *AccountHandler) GetAccountByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		userID, err := uuid.Parse(vars["userid"])
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		accountID, err := uuid.Parse(vars["id"])
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		data, err := a.service.GetAccountByIDAndUserID(accountID, userID)
		if err != nil && !errors.Is(err, account.ErrNotFound) {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		if errors.Is(err, account.ErrNotFound) || data == nil {
			a.logger.Sugar().Errorf("Account with ID: %v not found", accountID)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		toJSON := &mapper.Account{
			ID:           data.ID,
			UserID:       data.UserID,
			Account:      data.Account,
			Currency:     data.Currency,
			Name:         data.Name,
			Amount:       data.Amount,
			InterestRate: data.InterestRate,
			IsActive:     data.IsActive,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
		}

		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading account"))
			if err != nil {
				a.logger.Error(err.Error())
			}
		}
	})
}

// UpdateAccount godoc
// @title Update Account by ID
// @version 1.0
// @summary Update account details based on the provided ID.
// @description Update the account details using the provided user ID.
// @tags account-minibank
// @accept json
// @produce json
// @param id path string true "Account ID"
// @param userid path string true "User ID"
// @param account body AccountUpdateRequest true "Account Update Payload"
// @success 200 {string} string "Successfully updated account details"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "Account not found"
// @Security BasicAuth
// @router /accounts/{id} [put]
func (a *AccountHandler) UpdateAccount() http.Handler {
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

		userid, err := uuid.Parse(vars["userid"])
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		_, err = a.service.GetAccountByIDAndUserID(id, userid)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("No account for this user"))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		var input struct {
			Name         string  `json:"name"`
			InterestRate float64 `json:"interest_rate"`
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't decode request"))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		err = a.service.UpdateAccount(id, input.Name, input.InterestRate)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't update account: " + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}
		a.logger.Sugar().Infof("Account %v was updated", id)
	})
}

// DeleteAccount godoc
// @title Delete Account by ID
// @version 1.0
// @summary Delete account based on the provided ID.
// @description Delete the account using the provided account ID.
// @tags account-minibank
// @accept json
// @produce json
// @param id path string true "Account ID"
// @param userid path string true "User ID"
// @success 200 {string} string "Successfully deleted account"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "Account not found"
// @Security BasicAuth
// @router /accounts/{id} [delete]
func (a *AccountHandler) DeleteAccount() http.Handler {
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

		userid, err := uuid.Parse(vars["userid"])
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		_, err = a.service.GetAccountByIDAndUserID(id, userid)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("No account for this user"))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		err = a.service.DeleteAccount(id)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't delete account: " + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}
		a.logger.Sugar().Infof("Account %v was deleted", id)
	})
}

// ListAccountsByUserID godoc
// @title Get List of Accounts by User ID
// @version 1.0
// @summary Retrieve list of accounts based on the provided User ID.
// @description Fetch the list of accounts using the provided User ID.
// @tags account-minibank
// @accept json
// @produce json
// @param id path string true "User ID"
// @success 200 {array} mapper.Account "Successfully retrieved account details"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "Accounts not found"
// @Security BasicAuth
// @router /accounts [get]
func (a *AccountHandler) ListAccountsByUserID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		userID, err := uuid.Parse(vars["id"])
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		data, err := a.service.ListAccount(userID)
		if err != nil && !errors.Is(err, account.ErrNotFound) {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		if errors.Is(err, account.ErrNotFound) || data == nil {
			a.logger.Error("Accounts not found")
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		var toJSON []*mapper.Account
		for _, item := range data {
			toJSON = append(toJSON, &mapper.Account{
				ID:           item.ID,
				UserID:       item.UserID,
				Account:      item.Account,
				Currency:     item.Currency,
				Name:         item.Name,
				Amount:       item.Amount,
				InterestRate: item.InterestRate,
				IsActive:     item.IsActive,
				CreatedAt:    item.CreatedAt,
				UpdatedAt:    item.UpdatedAt,
			})
		}

		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading account"))
			if err != nil {
				a.logger.Error(err.Error())
			}
		}
	})
}

// TopUp godoc
// @title TopUp Account
// @version 1.0
// @summary TopUp account balance based on the provided ID.
// @description TopUp the account balance using the provided user ID and amount.
// @tags account-minibank
// @accept json
// @produce json
// @param id path string true "Account ID"
// @param userid path string true "User ID"
// @param account body ChangeBalanceRequest true "TopUp Account Payload"
// @success 200 {string} string "Successfully toped up account details"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "Account not found"
// @Security BasicAuth
// @router /accounts/{id}/topup [put]
func (a *AccountHandler) TopUp() http.Handler {
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

		userid, err := uuid.Parse(vars["userid"])
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		_, err = a.service.GetAccountByIDAndUserID(id, userid)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("No account for this user"))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		var input struct {
			Amount float64 `json:"amount"`
		}

		a.logger.Sugar().Debugf("input.Amount %v", input.Amount)

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't decode request"))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		balance, err := a.service.TopUp(id, input.Amount)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't top up account: " + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		toJSON := &struct {
			Balance float64 `json:"balance"`
		}{
			Balance: balance,
		}

		a.logger.Sugar().Debugf("Balance %v", balance)

		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading balance"))
			if err != nil {
				a.logger.Error(err.Error())
			}
		}
		a.logger.Sugar().Infof("Account %v was toped up", id)
	})
}

// Withdraw godoc
// @title Withdraw Account
// @version 1.0
// @summary Withdraw money based on the provided ID.
// @description Withdraw money using the provided user ID and amount.
// @tags account-minibank
// @accept json
// @produce json
// @param id path string true "Account ID"
// @param userid path string true "User ID"
// @param account body ChangeBalanceRequest true "Withdraw Account Payload"
// @success 200 {string} string "Successfully Withdrawed account"
// @failure 500 {string} string "Internal server error"
// @failure 404 {string} string "Account not found"
// @Security BasicAuth
// @router /accounts/{id}/withdraw [put]
func (a *AccountHandler) Withdraw() http.Handler {
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

		userid, err := uuid.Parse(vars["userid"])
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't parse ID:" + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		_, err = a.service.GetAccountByIDAndUserID(id, userid)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("No account for this user"))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		var input struct {
			Amount float64 `json:"amount"`
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't decode request"))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		a.logger.Sugar().Debugf("input.Amount %v", input.Amount)

		balance, err := a.service.WithDraw(id, input.Amount)
		if err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Couldn't withdraw account: " + err.Error()))
			if err != nil {
				a.logger.Error(err.Error())
			}

			return
		}

		toJSON := &struct {
			Balance float64 `json:"balance"`
		}{
			Balance: balance,
		}

		a.logger.Sugar().Debugf("Balance %v", balance)

		if err := json.NewEncoder(w).Encode(toJSON); err != nil {
			a.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte("Error while reading ID"))
			if err != nil {
				a.logger.Error(err.Error())
			}
		}
		a.logger.Sugar().Infof("Account %v was withdrawed", id)
	})
}