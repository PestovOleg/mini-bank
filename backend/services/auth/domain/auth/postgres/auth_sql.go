package postgres

import (
	"database/sql"
	"time"

	"github.com/PestovOleg/mini-bank/backend/services/auth/domain/auth"
	"github.com/google/uuid"
)

type AuthSQL struct {
	db *sql.DB
}

func NewAuthSQL(db *sql.DB) *AuthSQL {
	return &AuthSQL{
		db: db,
	}
}

func (r *AuthSQL) Create(u *auth.Auth) (uuid.UUID, error) {
	rec, err := r.db.Prepare(`
		insert into authentications (id, username, password, is_active, created_at)
		values($1, $2, $3, $4, $5)`)

	if err != nil {
		return u.ID, err
	}

	_, err = rec.Exec(
		u.ID,
		u.Username,
		u.Password,
		u.IsActive,
		time.Now().Format(time.RFC1123Z),
	)
	if err != nil {
		return u.ID, err
	}

	err = rec.Close()
	if err != nil {
		return u.ID, err
	}

	return u.ID, nil
}

func (r *AuthSQL) GetByID(id uuid.UUID) (*auth.Auth, error) {
	rec, err := r.db.Prepare(`
		select id, username, password,is_active, created_at,updated_at
		from authentications where id=$1`)

	if err != nil {
		return nil, err
	}

	var a auth.Auth

	rows, err := rec.Query(id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&a.ID,
			&a.Username,
			&a.Password,
			&a.IsActive,
			&a.CreatedAt,
			&a.UpdatedAt,
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

func (r *AuthSQL) GetByUName(username string) (*auth.Auth, error) {
	rec, err := r.db.Prepare(`
		select id, username, password,is_active, created_at,updated_at
		from authentications where username=$1`)

	if err != nil {
		return nil, err
	}

	var a auth.Auth

	rows, err := rec.Query(username)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&a.ID,
			&a.Username,
			&a.Password,
			&a.IsActive,
			&a.CreatedAt,
			&a.UpdatedAt,
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

func (r *AuthSQL) Update(u *auth.Auth) error {
	rec, err := r.db.Prepare(`
		update authentications set is_active=$1,updated_at=$2 where id=$3`)

	if err != nil {
		return err
	}

	_, err = rec.Exec(
		u.IsActive,
		u.UpdatedAt,
		u.ID,
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

func (r *AuthSQL) Delete(id uuid.UUID) error {
	rec, err := r.db.Prepare(`
		delete from authentications where id=$1`)

	if err != nil {
		return err
	}

	res, err := rec.Exec(id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return auth.ErrNotFound
	}

	err = rec.Close()
	if err != nil {
		return err
	}

	return nil
}
