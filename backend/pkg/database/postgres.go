package database

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	_ "github.com/lib/pq" // working with driver
)

type Postgres struct{}

//nolint:gochecknoglobals
var (
	sConns   *sql.DB
	sConnsMx sync.Mutex
)

func (p *Postgres) GetSQLDBCon(conn *DBCon) (*sql.DB, error) {
	sConnsMx.Lock()
	defer sConnsMx.Unlock()

	if sConns == nil {
		connStr := getConString(conn)
		sConns, err := sql.Open("postgres", connStr)

		if err != nil {
			return nil, err
		}

		logger := logger.GetLogger("db")
		logger.Debug(connStr)

		return sConns, nil
	}

	return sConns, nil
}

// возврат строки коннекта
func getConString(conn *DBCon) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		conn.User,
		conn.Password,
		conn.Host,
		conn.Port,
		conn.Name,
		conn.SSLMode,
	)
}

func NewPostgres() *Postgres {
	return &Postgres{}
}
