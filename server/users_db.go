package main

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// _________users_________________________________________________
// |  id      |  email    |  username  |  password  |  sessionId  |
// |  INTEGER |  TEXT     |  TEXT      |  TEXT      |  TEXT       |
func crerateUsersTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, email TEXT NOT NULL UNIQUE, username TEXT NOT NULL UNIQUE, password TEXT NOT NULL, sessionId TEXT)")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

func saveUser(username string, email string, password string, sessionId string) error {
	statement, err := db.Prepare("INSERT INTO users (email, username, password, sessionId) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(strings.ToLower(email), username, password, sessionId)
	if err != nil {
		return err
	}
	return nil
}
