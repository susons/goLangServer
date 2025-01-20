package models

import (
	"fmt"
	"time"

	"example.com/project_api/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

func (e Event) Save() error {
	// later: add to db
	query := `INSERT INTO EVENTS(name, description,location, dateTime, userId)
	VALUES(?,?,?,?,?)`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.Id = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)

	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.Id)

	return err
}
