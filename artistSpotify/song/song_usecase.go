package song

import "artistSpotify/model"

type SongUsecase interface {
	GetAllSong(id int) (*[]model.Song, error)
	InsertSong(data model.Song) (*model.Song,error)
}