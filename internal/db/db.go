// Package db provides a connection to the database
package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(conn string) {
	var err error
	DB, err = sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
}
