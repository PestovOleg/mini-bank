package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	"go.uber.org/zap"
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

var sConns *sql.DB
var sConnsMx sync.Mutex

func GetDBCon() *sql.DB {
	sConnsMx.Lock()
	defer sConnsMx.Unlock()

	if sConns == nil {
		connStr := getConString()
		sConns, err := sql.Open("postgres", connStr)

		if err != nil {
			fmt.Println("Error at opening database connection", zap.Error(err))
		}

		return sConns
	}

	return sConns
}

// return connection string
func getConString() string {

	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		con.User,
		con.Password,
		con.Host,
		con.Port,
		con.Name,
		con.SSLMode,
	)
}
