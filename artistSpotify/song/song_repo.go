package song

import "artistSpotify/model"

type SongRepo interface {
	GetAllSong(id int) (*[]model.Song, error)
	InsertSong(data model.Song) (*model.Song,error)
}
