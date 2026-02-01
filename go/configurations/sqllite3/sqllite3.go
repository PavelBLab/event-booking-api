package sqllite3

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to databases.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	//createTables()
}
