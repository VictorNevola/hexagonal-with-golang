package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectToMySql(url string) *sqlx.DB {
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		panic("connection error")
	}

	return db
}
