package database

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:ams.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
