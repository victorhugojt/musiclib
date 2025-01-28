package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error // Declare err
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Error connecting to sqlite")
	}

	if err = DB.Ping(); err != nil {
		panic(fmt.Sprintf("Error pinging database: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	usersTableStr := ` CREATE TABLE IF NOT EXISTS users (
						id TEXT PRIMARY KEY,
						full_name TEXT NOT NULL,
						email TEXT NOT NULL,
						created_by DATETIME NOT NULL,
						created_on DATETIME NOT NULL,
						user_name TEXT NOT NULL,
						password TEXT NOT NULL
						)
					`

	_, err := DB.Exec(usersTableStr)
	if err != nil {
		panic("error creating USERS table")
	}

	librariesTableStr := ` CREATE TABLE IF NOT EXISTS libraries (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_by DATETIME NOT NULL,
		created_on DATETIME NOT NULL
		)
	`

	_, err = DB.Exec(librariesTableStr)
	if err != nil {
		panic("error creating LIBRARIES table")
	}

}
