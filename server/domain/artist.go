package domain

import "time"

type Artist struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Dob                time.Time `json:"-"`
	Gender             string    `json:"gender"`
	Address            string    `json:"address"`
	FirstReleaseYear   string    `json:"first_release_year"`
	NoOfAlbumsReleased int       `json:"no_of_albums_released"`
	CreatedAt          time.Time `json:"-"`
	UpdatedAt          time.Time `json:"-"`
}
