package service

import (
	"artist-management-system/domain"
	"artist-management-system/view"
	"database/sql"
	"fmt"
)

type artistService struct {
	db *sql.DB
}

type ArtistService interface {
	All() ([]domain.Artist, error)
	Create(params view.ArtistView) error
	Update(id int, params view.ArtistView) error
	Delete(id int) error
}

func NewArtistService(db *sql.DB) ArtistService {
	return artistService{
		db: db,
	}
}

func (service artistService) All() ([]domain.Artist, error) {
	query := `
		SELECT id, name, gender, address, first_release_year, no_of_albums_released FROM artist;
	`

	var artists []domain.Artist
	rows, err := service.db.Query(query)
	if err != nil {
		fmt.Printf("ERROR:: could not query database: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var artist domain.Artist
		if err := rows.Scan(&artist.ID, &artist.Name, &artist.Gender, &artist.Address, &artist.FirstReleaseYear, &artist.NoOfAlbumsReleased); err != nil {
			fmt.Printf("ERROR:: could not scan values from row: %s\n", err.Error())
			return nil, err
		}
		artists = append(artists, artist)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("ERROR:: could not scan next row: %s\n", err.Error())
		return nil, err
	}

	return artists, nil
}

func (service artistService) Create(params view.ArtistView) error {
	query := `
		INSERT INTO artist (name, gender, address, first_release_year, no_of_albums_released)
		VALUES ($1, $2, $3, $4, $5);
	`

	if _, err := service.db.Exec(query, &params.Name, &params.Gender, &params.Address, &params.FirstReleaseYear, &params.NoOfAlbumsReleased); err != nil {
		fmt.Printf("ERROR:: could not insert to artist table: %s\n", err.Error())
		return err
	}

	return nil
}

func (service artistService) Update(id int, params view.ArtistView) error {
	updateQuery := `
		UPDATE artist
		SET name = $1, gender = $2, address = $3, first_release_year = $4, no_of_albums_released = $5
		WHERE id = $6
	`
	if _, err := service.db.Exec(updateQuery, &params.Name, &params.Gender, &params.Address, &params.FirstReleaseYear, &params.NoOfAlbumsReleased, &id); err != nil {
		return err
	}

	return nil
}

func (service artistService) Delete(id int) error {
	query := `
		DELETE FROM artist WHERE id = $1
	`

	if _, err := service.db.Exec(query, &id); err != nil {
		return err
	}

	return nil
}
