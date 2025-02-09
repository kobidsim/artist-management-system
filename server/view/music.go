package view

type MusicView struct {
	Title        string `json:"title" validate:"required"`
	ComposedByID string `json:"composed_by_id" validate:"required"`
	AlbumName    string `json:"album_name"`
	Genre        string `json:"genre" validation:"oneof=rnb country classic rock jazz"`
}

type UpdateMusicView struct {
	Title     string `json:"title" validate:"required"`
	AlbumName string `json:"album_name"`
	Genre     string `json:"genre" validation:"oneof=rnb country classic rock jazz"`
}
