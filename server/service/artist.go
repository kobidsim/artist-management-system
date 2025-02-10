package service

import (
	"artist-management-system/domain"
	"artist-management-system/view"
	"bytes"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"slices"
	"strconv"
	"time"
)

type artistService struct {
	db *sql.DB
}

type ArtistService interface {
	All() ([]domain.Artist, error)
	Create(params view.ArtistView) error
	Update(id int, params view.ArtistView) error
	Delete(id int) error
	CSVImport(file multipart.File) error
	CSVExport() (*bytes.Buffer, error)
}

func NewArtistService(db *sql.DB) ArtistService {
	return artistService{
		db: db,
	}
}

func (service artistService) All() ([]domain.Artist, error) {
	query := `
		SELECT id, name, gender, address, first_release_year, no_of_albums_released, dob FROM artist;
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
		if err := rows.Scan(&artist.ID, &artist.Name, &artist.Gender, &artist.Address, &artist.FirstReleaseYear, &artist.NoOfAlbumsReleased, &artist.Dob); err != nil {
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
	createdAt := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")

	query := `
		INSERT INTO artist (name, gender, address, first_release_year, no_of_albums_released, dob, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	if _, err := service.db.Exec(query, &params.Name, &params.Gender, &params.Address, &params.FirstReleaseYear, &params.NoOfAlbumsReleased, &params.Dob, &createdAt); err != nil {
		fmt.Printf("ERROR:: could not insert to artist table: %s\n", err.Error())
		return err
	}

	return nil
}

func (service artistService) Update(id int, params view.ArtistView) error {
	updatedAt := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")

	updateQuery := `
		UPDATE artist
		SET name = $1, gender = $2, address = $3, first_release_year = $4, no_of_albums_released = $5, dob = $6, updated_at = $7
		WHERE id = $8
	`
	if _, err := service.db.Exec(updateQuery, &params.Name, &params.Gender, &params.Address, &params.FirstReleaseYear, &params.NoOfAlbumsReleased, &params.Dob, &updatedAt, &id); err != nil {
		return err
	}

	return nil
}

func (service artistService) Delete(id int) error {
	deleteMusicQuery := `
		DELETE FROM music WHERE artist_id = $1
	`
	if _, err := service.db.Exec(deleteMusicQuery, &id); err != nil {
		return err
	}

	query := `
		DELETE FROM artist WHERE id = $1
	`
	if _, err := service.db.Exec(query, &id); err != nil {
		return err
	}

	return nil
}

func (service artistService) CSVImport(file multipart.File) error {
	reader := csv.NewReader(file)
	data := map[string][]string{
		"name":                  {},
		"gender":                {},
		"address":               {},
		"first_release_year":    {},
		"no_of_albums_released": {},
	}
	header, err := reader.Read()
	if err != nil {
		return errors.New("error reading file. file may not be csv")
	}
	if !(slices.Contains(header, "name") &&
		slices.Contains(header, "gender") &&
		slices.Contains(header, "address") &&
		slices.Contains(header, "first_release_year") &&
		slices.Contains(header, "no_of_albums_released") && len(header) == 5) {
		return errors.New("csv does not contain the correct columns")
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.New("error reading file")
		}
		for i := 0; i < len(record); i++ {
			_, ok := data[header[i]]
			if ok {
				data[header[i]] = append(data[header[i]], record[i])
			}
		}
	}

	query := `
		INSERT INTO artist (name, gender, address, first_release_year, no_of_albums_released)
		VALUES ($1, $2, $3, $4, $5)
	`
	tx, err := service.db.Begin()
	if err != nil {
		return errors.New("error adding data")
	}
	defer tx.Rollback()

	for i := 0; i < len(data["name"]); i++ {
		if !slices.Contains([]string{"m", "f", "o"}, data["gender"][i]) {
			return errors.New("gender is invalid")
		}

		noOfAlbumsReleased, err := strconv.Atoi(data["no_of_albums_released"][i])
		if err != nil {
			return errors.New("number of albums released is not a number")
		}

		if _, err := tx.Exec(query, &data["name"][i], &data["gender"][i],
			&data["address"][i], &data["first_release_year"][i],
			&noOfAlbumsReleased); err != nil {
			return errors.New("error adding data")
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.New("error adding data")
	}
	return nil
}

func (service artistService) CSVExport() (*bytes.Buffer, error) {
	artists, err := service.All()
	if err != nil {
		return nil, err
	}

	var csvBuffer bytes.Buffer
	csvWriter := csv.NewWriter(&csvBuffer)

	csvWriter.Write([]string{"name", "gender", "address", "first_release_year", "no_of_albums_released"})

	for _, artist := range artists {
		csvWriter.Write([]string{artist.Name, artist.Gender, artist.Address, artist.FirstReleaseYear, strconv.Itoa(artist.NoOfAlbumsReleased)})
	}
	csvWriter.Flush()

	return &csvBuffer, nil
}
