package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./events.db")
	if err != nil {
		log.Fatalf("Failed to open SQLite DB: %s", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS sensor_events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sensor TEXT,
		value TEXT,
		timestamp TEXT
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create table: %s", err)
	}

	log.Println("SQLite DB initialized and table ready.")
}

func InsertEvent(sensor, value, timestamp string) {
	stmt, err := DB.Prepare("INSERT INTO sensor_events(sensor, value, timestamp) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Error preparing insert:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(sensor, value, timestamp)
	if err != nil {
		log.Println("Error inserting event:", err)
	}
}
