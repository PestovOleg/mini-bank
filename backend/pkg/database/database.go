package database

import (
	"database/sql"
)

type Database interface {
	GetSQLDBCon(conn *DBCon) (*sql.DB, error)
}

type DBCon struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

func NewDBCon(user, password, host, port, name, sslMode string) *DBCon {
	return &DBCon{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name,
		SSLMode:  sslMode,
	}
}

func NewDatabase() Database {
	return &Postgres{}
}
