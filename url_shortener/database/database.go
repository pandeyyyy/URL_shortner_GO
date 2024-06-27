package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS urls (
        short_url TEXT PRIMARY KEY,
        long_url TEXT NOT NULL
    );
    `
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func CloseDB() {
	db.Close()
}

func StoreURL(shortURL, longURL string) error {
	insertQuery := `INSERT INTO urls (short_url, long_url) VALUES (?, ?)`
	_, err := db.Exec(insertQuery, shortURL, longURL)
	return err
}

func RetrieveURL(shortURL string) (string, error) {
	var longURL string
	selectQuery := `SELECT long_url FROM urls WHERE short_url = ?`
	err := db.QueryRow(selectQuery, shortURL).Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
