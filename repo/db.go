package repo

import (
	"database/sql"
	"os"
)

var db *sql.DB

func InitDatabase() {
	var err error

	db, err = sql.Open("mysql", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}
}
