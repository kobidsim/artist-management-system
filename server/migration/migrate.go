package main

import (
	"artist-management-system/database"
	"log"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database. ERROR::%s", err.Error())
	}
	defer db.Close()

	schema := `
		DROP TABLE IF EXISTS user;

		DROP TABLE IF EXISTS artist;

		DROP TABLE IF EXISTS music;

		DROP TABLE IF EXISTS invalid_tokens;

		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_name VARCHAR(255),
			last_name VARCHAR(255),
			role TEXT CHECK( role IN ('super_admin', 'artist_manager', 'artist') ) NOT NULL,
			email VARCHAR(255) UNIQUE,
			password VARCHAR(500),
			phone VARCHAR(20),
			dob TEXT,
			gender TEXT CHECK( gender IN ('m', 'f', 'o') ),
			address VARCHAR(255),
			created_at TEXT,
			updated_at TEXT
		);

		CREATE TABLE IF NOT EXISTS artist (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255),
			dob TEXT,
			gender TEXT CHECK( gender IN ('m', 'f', 'o') ),
			address VARCHAR(255),
			first_release_year TEXT,
			no_of_albums_released INTEGER,
			created_at TEXT,
			updated_at TEXT
		);

		CREATE TABLE IF NOT EXISTS music (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(255),
			artist_id INTEGER,
			album_name VARCHAR(255),
			genre TEXT CHECK( genre IN ('rnb', 'country', 'classic', 'rock', 'jazz') ),
			created_at TEXT,
			updated_at TEXT,
			FOREIGN KEY (artist_id) REFERENCES artist(id)
		);

		CREATE TABLE IF NOT EXISTS invalid_tokens (
			token TEXT
		);
	`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create tables. ERROR::%s", err.Error())
	}
}
