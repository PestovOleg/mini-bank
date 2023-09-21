package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	_ "github.com/lib/pq" // working with driver
)

type dbcon struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

func NewDBCon(user, password, host, port, name, sslMode string) *dbcon {
	return &dbcon{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name,
		SSLMode:  sslMode,
	}
}

//nolint:gochecknoglobals
var (
	sConns   *sql.DB
	sConnsMx sync.Mutex
)

func GetDBCon(conn *dbcon) (*sql.DB, error) {
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
func getConString(conn *dbcon) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		conn.User,
		conn.Password,
		conn.Host,
		conn.Port,
		conn.Name,
		conn.SSLMode,
	)
}
