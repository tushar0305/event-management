package models

import (
	"time"
	"github.com/tushar0305/event-management/db"
	"fmt"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      int64       `json:"user_id"`
}

func (e *Event) Save() error {
    if db.DB == nil {
        return fmt.Errorf("database connection is not initialized")
    }

    query := `INSERT INTO events(name, description, location, dateTime, user_id) 
    VALUES(?, ?, ?, ?, ?)`
    
    stmt, err := db.DB.Prepare(query)
    if err != nil {
        fmt.Println("Error preparing query:", err)
        return err
    }
    defer stmt.Close()
    
    result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return err
    }
    id, err := result.LastInsertId()
    if err != nil {
        fmt.Println("Error getting last insert id:", err)
        return err
    }
    e.Id = id

    return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return []Event{}, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return []Event{}, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event *Event) UpdateEventById() error {
    query := `
    UPDATE events
    SET name = ?, description = ?, location = ?, dateTime = ?
    WHERE id = ?
    `

    stmt, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
    if err != nil {
        return err
    }
    return nil
}

func (event Event) DeleteEventById() error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Id)
	return err
}