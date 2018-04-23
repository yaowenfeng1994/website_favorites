package db

import (
	"database/sql"
)

func init() {
	db, err := sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/website_project?charset=utf8")
	if err == nil {
		db.SetMaxOpenConns(2000)
		db.SetMaxIdleConns(1000)
		db.Ping()
	}
}
