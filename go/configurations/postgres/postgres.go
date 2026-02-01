package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=api sslmode=disable")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	//createTables()
}