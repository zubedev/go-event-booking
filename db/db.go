package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "db.sqlite")
	if err != nil {
		panic(err)
	}
	createTables()
}

func createTables() {
	// Create users table
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    email TEXT NOT NULL UNIQUE,
	    password TEXT NOT NULL
	);`
	_, err := DB.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	// Create events table
	sqlStmt = `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER,
	    FOREIGN KEY(userId) REFERENCES users(id)
	);`
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}
