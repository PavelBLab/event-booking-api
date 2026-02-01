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
		RETURNING id
	`
	err := postgres.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.ID)

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

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (Event, error) {
	query := `SELECT * FROM events WHERE id = $1`
	row := postgres.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return Event{}, err
	}

	return event, nil
}

func (e Event) Update() error {
	query := `UPDATE events 
			  SET name = $1, description = $2, location = $3, date_time = $4 
			  WHERE id = $5`
	statement, err := postgres.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = $1`
	statement, err := postgres.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.ID)

	return err
}
