package models

import (
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	Date        time.Time `json:"date"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() {
	// later: add it to the databases

	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
