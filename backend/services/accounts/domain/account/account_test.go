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
			name     string
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
				name     string
			}{userID: uuid.Must(uuid.Parse("99371a5f-ad18-4d75-8e4f-ce1e4cbd6dac")),
				currency: "810",
				account:  "40817810902000000000",
				name:     "так себе счет"},
			out: struct {
				account string
				err     error
			}{account: "40817810902000000001", err: nil},
		},
		{
			in: struct {
				userID   uuid.UUID
				currency string
				account  string
				name     string
			}{userID: uuid.Must(uuid.Parse("99371a5f-ad18-4d75-8e4f-ce1e4cbd6dac")),
				currency: "810",
				account:  "",
				name:     "так себе счет"},
			out: struct {
				account string
				err     error
			}{account: "40817810902000000001", err: nil},
		},
	}

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			a, err := NewAccount(i.in.userID, i.in.currency, i.in.account, i.in.name)
			if a.Account != i.out.account || err != nil {
				t.Errorf("got %v and %v, wants %v", a, err, i.out)
			}
		})
	}
}
