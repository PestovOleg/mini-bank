package postgres

import (
	"database/sql"
	"time"

	"github.com/PestovOleg/mini-bank/backend/domain/user"
	"github.com/google/uuid"
)

type UserSQL struct {
	db *sql.DB
}

func NewUserSQL(db *sql.DB) *UserSQL {
	return &UserSQL{
		db: db,
	}
}

func (r *UserSQL) Create(u *user.User) (uuid.UUID, error) {
	rec, err := r.db.Prepare(`
		insert into users (id, email, phone,name, last_name, patronymic, created_at, birthday)
		values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`)

	if err != nil {
		return u.ID, err
	}

	_, err = rec.Exec(
		u.ID,
		u.Email,
		u.Phone,
		u.Name,
		u.LastName,
		u.Patronymic,
		time.Now().Format(time.RFC1123Z),
		u.Birthday,
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

func (r *UserSQL) Get(id uuid.UUID) (*user.User, error) {
	rec, err := r.db.Prepare(`
		select id, email, phone, birthday, name, last_name, patronymic, created_at,updated_at
		from users where id=$1`)

	if err != nil {
		return nil, err
	}

	var u user.User

	rows, err := rec.Query(id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&u.ID,
			&u.Email,
			&u.Phone,
			&u.Birthday,
			&u.Name,
			&u.LastName,
			&u.Patronymic,
			&u.CreatedAt,
			&u.UpdatedAt,
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

	return &u, nil
}

func (r *UserSQL) Update(u *user.User) error {
	rec, err := r.db.Prepare(`
		update users set email=$1, updated_at=$2,phone=$3 where id=$4`)

	if err != nil {
		return err
	}

	_, err = rec.Exec(
		u.Email,
		u.UpdatedAt,
		u.Phone,
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
