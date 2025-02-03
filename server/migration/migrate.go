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
		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_name VARCHAR(255),
			last_name VARCHAR(255),
			role TEXT CHECK( role IN ('super_admin', 'artist_manager', 'artist') ) NOT NULL,
			email VARCHAR(255),
			password VARCHAR(500),
			phone VARCHAR(20),
			dob DATETIME,
			gender TEXT CHECK( gender IN ('m', 'f', 'o') ),
			address VARCHAR(255),
			created_at DATETIME,
			updated_at DATETIME
		);

		CREATE TABLE IF NOT EXISTS artist (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255),
			dob DATETIME,
			gender TEXT CHECK( gender IN ('m', 'f', 'o') ),
			address VARCHAR(255),
			first_release_year YEAR,
			no_of_albums_released INTEGER,
			created_at DATETIME,
			updated_at DATETIME
		);

		CREATE TABLE IF NOT EXISTS music (
			title VARCHAR(255),
			artist_id INTEGER,
			album_name VARCHAR(255),
			genre TEXT CHECK( genre IN ('rnb', 'country', 'classic', 'rock', 'jazz') ),
			created_at DATETIME,
			updated_at DATETIME,
			FOREIGN KEY (artist_id) REFERENCES artist(id)
		);
	`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create tables. ERROR::%s", err.Error())
	}
}
