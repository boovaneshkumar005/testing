package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // pure Go SQLite driver
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./test.db") // note "sqlite" driver name
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL
    );
    `
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
func CloseDB() {
	DB.Close()
}
