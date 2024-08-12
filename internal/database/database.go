package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

const (
	username = "mysql"
	password = "mysql"
)

// create connection of database
func CreateConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   username,
		Passwd: password,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "blueSoftBank",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
