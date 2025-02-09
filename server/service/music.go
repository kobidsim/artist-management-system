package service

import (
	"artist-management-system/domain"
	"artist-management-system/view"
	"database/sql"
	"fmt"
	"strconv"
)

type musicService struct {
	db *sql.DB
}

type MusicService interface {
	All(artistID int) ([]domain.Music, error)
	Create(params view.MusicView) error
	Update(id int, params view.UpdateMusicView) error
	Delete(id int) error
}

func NewMusicService(db *sql.DB) MusicService {
	return musicService{
		db: db,
	}
}

func (service musicService) All(artistID int) ([]domain.Music, error) {
	query := `
		SELECT m.id, m.title, m.album_name, m.genre, m.artist_id,
			a.id, a.name, a.gender, a.address, a.first_release_year, a.no_of_albums_released
		FROM music as m
		LEFT JOIN artist as a ON m.artist_id = a.id
		WHERE artist_id = $1;
	`

	var musicList []domain.Music
	rows, err := service.db.Query(query, artistID)
	if err != nil {
		fmt.Printf("ERROR:: could not query database: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var music domain.Music
		if err := rows.Scan(&music.ID, &music.Title, &music.AlbumName, &music.Genre, &music.ArtistID, &music.Artist.ID, &music.Artist.Name,
			&music.Artist.Gender, &music.Artist.Address, &music.Artist.FirstReleaseYear, &music.Artist.NoOfAlbumsReleased); err != nil {
			fmt.Printf("ERROR:: could not scan values from row: %s\n", err.Error())
			return nil, err
		}
		musicList = append(musicList, music)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("ERROR:: could not scan next row: %s\n", err.Error())
		return nil, err
	}

	return musicList, nil
}

func (service musicService) Create(params view.MusicView) error {
	query := `
		INSERT INTO music (title, album_name, genre, artist_id)
		VALUES ($1, $2, $3, $4);
	`

	artistID, _ := strconv.Atoi(params.ComposedByID)

	if _, err := service.db.Exec(query, &params.Title, &params.AlbumName, &params.Genre, &artistID); err != nil {
		fmt.Printf("ERROR:: could not insert to artist table: %s\n", err.Error())
		return err
	}

	return nil
}

func (service musicService) Update(id int, params view.UpdateMusicView) error {
	updateQuery := `
		UPDATE music
		SET title = $1, album_name = $2, genre = $3
		WHERE id = $4
	`
	if _, err := service.db.Exec(updateQuery, &params.Title, &params.AlbumName, &params.Genre, &id); err != nil {
		return err
	}

	return nil
}

func (service musicService) Delete(id int) error {
	query := `
		DELETE FROM music WHERE id = $1
	`

	if _, err := service.db.Exec(query, &id); err != nil {
		return err
	}

	return nil
}
