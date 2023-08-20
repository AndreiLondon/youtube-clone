package main

import (
	"database/sql"

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
