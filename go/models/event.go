package models

import (
	"time"

	"github.com/PavelBLab/event-booking-api/configurations/postgres"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int       `json:"userId"`
}

var events = []Event{}

func (e Event) Save() error {
	// later: add it to the databases

	query := `
		INSERT INTO events (name, description, location, date_time, user_id) 
		VALUES ($1, $2, $3, $4, $5)
	`
	statement, err := postgres.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := postgres.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, Event{})
	}

	return events, nil
}
