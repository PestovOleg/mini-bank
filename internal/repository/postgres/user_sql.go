package postgres

import (
	"database/sql"
	"time"

	"github.com/PestovOleg/mini-bank/entity"
	"github.com/google/uuid"
)

type UserSQL struct {
	db *sql.DB
}

func (r *UserSQL) Create(u *entity.User) (uuid.UUID, error) {
	rec, err := r.db.Prepare(`
		insert into user (id, username, email, name, last_name, patronymic, password, is_active, created_at)
		values(?,?,?,?,?,?,?,?)`)
	if err != nil {
		return u.ID, err
	}

	_, err = rec.Exec(
		u.ID,
		u.Username,
		u.Email,
		u.Name,
		u.LastName,
		u.Patronymic,
		u.Password,
		u.IsActive,
		time.Now().Format("2023-09-10"),
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
