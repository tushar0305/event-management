package db

import (
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
    var err error
    DB, err = sql.Open("sqlite3", "api.db")
    if err != nil {
        panic("could not connect to database")
    }

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    createTables()
}

func createTables() {
    createUserTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );
    `

    _, err := DB.Exec(createUserTable)
    if err != nil {
        panic("Could not create user table!")
    }

    createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INTEGER
    );
    `

    _, err = DB.Exec(createEventsTable)
    if err != nil {
        panic("Could not create event table!")
    }
}

