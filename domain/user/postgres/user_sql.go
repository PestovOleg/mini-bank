package postgres

import (
	"database/sql"
	"time"

	"github.com/PestovOleg/mini-bank/domain/user"
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

func (r *UserSQL) Get(id uuid.UUID) (*user.User, error) {
	return nil, nil
}

func (r *UserSQL) List() ([]*user.User, error) {
	return nil, nil
}

func (r *UserSQL) Update(u *user.User) error {
	return nil
}
