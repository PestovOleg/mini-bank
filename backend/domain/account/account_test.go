package account

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestNewAccountTableDriven(t *testing.T) {
	tests := []struct {
		in struct {
			userID   uuid.UUID
			currency string
			account  string
		}
		out struct {
			account string
			err     error
		}
	}{
		{
			in: struct {
				userID   uuid.UUID
				currency string
				account  string
			}{userID: uuid.Must(uuid.Parse("99371a5f-ad18-4d75-8e4f-ce1e4cbd6dac")), currency: "810", account: ""},
			out: struct {
				account string
				err     error
			}{account: "40817810902000000123", err: nil},
		},
		{
			in: struct {
				userID   uuid.UUID
				currency string
				account  string
			}{userID: uuid.Must(uuid.Parse("99371a5f-ad18-4d75-8e4f-ce1e4cbd6dac")), currency: "810", account: ""},
			out: struct {
				account string
				err     error
			}{account: "40817810902000000123", err: nil},
		},
	}

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			a, err := NewAccount(i.in.userID, i.in.currency, i.in.account)
			if a.Account != i.out.account || err != nil {
				t.Errorf("got %v and %v, wants %v", a, err, i.out)
			}
		})
	}
}
