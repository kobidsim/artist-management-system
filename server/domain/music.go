package domain

import (
	"time"
)

type Music struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	ArtistID  int       `json:"artist_id"`
	Artist    Artist    `json:"artist"`
	AlbumName string    `json:"album_name"`
	Genre     string    `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
