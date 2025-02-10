package view

type ArtistView struct {
	Name               string `json:"name" validate:"required"`
	Dob                string `json:"dob"`
	Gender             string `json:"gender" validate:"oneof=m f o"`
	Address            string `json:"address"`
	FirstReleaseYear   string `json:"first_release_year"`
	NoOfAlbumsReleased int    `json:"no_of_albums_released"`
}
