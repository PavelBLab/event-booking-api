package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost port=7001 user=postgres password=postgres dbname=event_booking_api sslmode=disable")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
		    id SERIAL PRIMARY KEY NOT NULL ,
		    name TEXT NOT NULL,
		    description TEXT NOT NULL,
		    location TEXT NOT NULL,
		    date_time TIMESTAMP NOT NULL,
		    user_id INTEGER
		)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table: " + err.Error())
	}
}
