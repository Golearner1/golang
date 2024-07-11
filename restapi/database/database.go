package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Initdatabse() {
	var err error
	DB, err = sql.Open("sqlite3", "restapi.db")
	if err != nil {
		panic("could not get databse")
	}
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	Createtables()
}
func Createtables() {
	createuserdb := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	email TEXT NOT NULL UNIQUE, 
 	password TEXT NOT NULL
	)	
	`
	_, err := DB.Exec(createuserdb)
	if err != nil {
		panic("failed TO CREATE USERS TABLE")
	}
	createeventdb := `
	CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
		authid INTEGER,
		FOREIGN KEY(authid) REFERENCES users (id)
	)	
	`
	_, err = DB.Exec(createeventdb)
	if err != nil {
		fmt.Print(err)
		panic("failed TO CREATE TABLE")
	}
	createregistrationsdb := `
	CREATE TABLE IF NOT EXISTS registrations ( 
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	auth_id INTEGER,
	FOREIGN KEY (event_id) REFERENCES events (id),
	FOREIGN KEY(auth_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createregistrationsdb)
	if err != nil {
		fmt.Print(err)
		panic("failed TO REGISTRATION CREATE TABLE")
	}
}
