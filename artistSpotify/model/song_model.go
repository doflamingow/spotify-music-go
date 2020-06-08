package model


type Song struct {
	ID int `from:"ID" json:"id,omitempty"`
	Name string `from:"Name" json:"artist_name,omitempty"`
	SongName string `from:"SongName" json:"song_name,omitempty"`
	ImageUrl string `from:"ImageUrl" json:"image,omitempty"`
}