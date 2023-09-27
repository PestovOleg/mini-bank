package postgres

import (
	"database/sql"
	"time"

	"github.com/PestovOleg/mini-bank/backend/domain/account"
	"github.com/google/uuid"
)

type AccountSQL struct {
	db *sql.DB
}

func NewAccountSQL(db *sql.DB) *AccountSQL {
	return &AccountSQL{
		db: db,
	}
}

func (r *AccountSQL) GetByID(id uuid.UUID) (*account.Account, error) {
	rec, err := r.db.Prepare(`
		select id, account, currency,name, amount, interest_rate, user_id, created_at,updated_at,is_active
		from accounts where id=$1 and is_active=true`)

	if err != nil {
		return nil, err
	}

	var a account.Account

	rows, err := rec.Query(id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&a.ID,
			&a.Account,
			&a.Currency,
			&a.Name,
			&a.Amount,
			&a.InterestRate,
			&a.UserID,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.IsActive,
		)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	err = rec.Close()
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *AccountSQL) GetByNumber(acc string) (*account.Account, error) {
	rec, err := r.db.Prepare(`
		select id, account, currency,name, amount, interest_rate, user_id, created_at,updated_at,is_active
		from accounts where account=$1 and is_active=true`)

	if err != nil {
		return nil, err
	}

	var a account.Account

	rows, err := rec.Query(acc)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&a.ID,
			&a.Account,
			&a.Currency,
			&a.Name,
			&a.Amount,
			&a.InterestRate,
			&a.UserID,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.IsActive,
		)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	err = rec.Close()
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *AccountSQL) List(userID uuid.UUID) ([]*account.Account, error) {
	rec, err := r.db.Prepare(`
		select id, account, currency,name, amount, interest_rate, user_id, created_at,updated_at,is_active 
		from accounts where user_id=$1 and is_active=true`)

	if err != nil {
		return nil, err
	}

	rows, err := rec.Query(userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*account.Account

	for rows.Next() {
		var a account.Account

		err = rows.Scan(
			&a.ID,
			&a.Account,
			&a.Currency,
			&a.Name,
			&a.Amount,
			&a.InterestRate,
			&a.UserID,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.IsActive,
		)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	err = rec.Close()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *AccountSQL) Update(a *account.Account) error {
	_, err := r.db.Exec(`
    update accounts set name=$1, interest_rate=$2, updated_at=$3,amount=$4
    where id=$5 and is_active=true`,
		a.Name,
		a.InterestRate,
		a.UpdatedAt,
		a.Amount,
		a.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountSQL) GetLastOpenedAccount(currency string) (string, error) {
	var account string

	rec, err := r.db.Prepare(`
		select account from accounts where currency=$1 order by account desc limit 1`)

	if err != nil {
		return "", err
	}

	rows, err := rec.Query(currency)

	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		// rows is empty, return ""
		return "", nil
	}

	err = rows.Scan(&account) // добавлен амперсанд перед account
	if err != nil {
		return "", err
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	err = rec.Close()
	if err != nil {
		return "", err
	}

	return account, nil
}

func (r *AccountSQL) Create(a *account.Account) error {
	rec, err := r.db.Prepare(`
		insert into accounts (id, account, currency,name, amount, interest_rate, user_id, created_at,updated_at,is_active)
		values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)

	if err != nil {
		return err
	}

	_, err = rec.Exec(
		a.ID,
		a.Account,
		a.Currency,
		a.Name,
		a.Amount,
		a.InterestRate,
		a.UserID,
		a.CreatedAt,
		a.UpdatedAt,
		a.IsActive,
	)
	if err != nil {
		return err
	}

	err = rec.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountSQL) Delete(id uuid.UUID) error {
	rec, err := r.db.Prepare(`
		update accounts set is_active=$1,updated_at=$2 where id=$3`)

	if err != nil {
		return err
	}

	_, err = rec.Exec(
		false,
		time.Now(),
		id,
	)
	if err != nil {
		return err
	}

	err = rec.Close()
	if err != nil {
		return err
	}

	return nil
}
