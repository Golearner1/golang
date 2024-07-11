package models

import (
	"time"

	"example.com/first-app/database"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	Authid      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, dateTime, authid)
VALUES (?, ?, ?, ?, ?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.Authid)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}
func Getallevents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.Authid)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?" //?for putting value
	row := database.DB.QueryRow(query, id)       //putting value through  queryrow()
	var event Event                              //crreating variaable of type struct
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.Authid)
	if err != nil {
		return nil, err
	}
	return &event, err
}
func (e Event) Update() error {
	query := `UPDATE events
SET name = ?, description = ?,location = ?,datetime=?
WHERE id = ?
`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.ID)
	return err
}
func (e Event) Delete() error {
	query := `DELETE FROM events WHERE ID=?`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}

func (e Event) Register(Authid int64) error {
	query := `INSERT INTO registrations(event_id, auth_id) VALUES (?, ?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, Authid)
	return err
}
func (e Event) Deleteregistration(Authid int64) error {
	query := `DELETE FROM registrations WHERE event_id=? AND auth_id=? `
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}
