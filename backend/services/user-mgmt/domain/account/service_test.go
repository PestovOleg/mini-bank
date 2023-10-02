package account

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestCreateAccountTableDriven(t *testing.T) {
	tests := []struct {
		userID   uuid.UUID
		currency string
		name     string
		out      struct {
			a   *Account
			err error
		}
	}{
		{
			userID:   uuid.MustParse("35ec1e95-87d3-42fe-8158-bf59c78b9e26"),
			currency: "810",
			name:     "Удачный",
			out: struct {
				a   *Account
				err error
			}{
				a: &Account{
					UserID:       uuid.MustParse("35ec1e95-87d3-42fe-8158-bf59c78b9e26"),
					Account:      "ACC123456",
					Currency:     "810",
					Amount:       1000.00,
					InterestRate: 0.02,
					IsActive:     true,
				},
				err: nil,
			},
		},
	}

	mockRepo := NewMockAccountRepository()

	service := NewService(mockRepo)

	for _, i := range tests {
		testname := fmt.Sprintf("input userID: %v, currency: %s wants out: %v, %v", i.userID, i.currency, i.out.a, i.out.err)
		t.Run(testname, func(t *testing.T) {
			acc, err := service.CreateAccount(i.userID, i.currency, i.name)
			if acc.ID != i.out.a.ID && !errors.Is(err, i.out.err) {
				t.Errorf("got %v and %v, wants %v and %v", acc.ID, err, i.out.a.ID, i.out.err)
			}
		})
	}
}
