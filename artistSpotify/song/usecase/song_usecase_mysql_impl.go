package usecase

import (
	"artistSpotify/model"
	"artistSpotify/song"
)

type SongUsecaseImpl struct {
	songRepo song.SongRepo
}

func CreateSongUsecaseImpl(songRepo song.SongRepo) song.SongRepo {
	return &SongUsecaseImpl{songRepo}
}

func (s *SongUsecaseImpl) GetAllSong(id int) (*[]model.Song, error) {
	return s.songRepo.GetAllSong(id)
}

func (s *SongUsecaseImpl) InsertSong(data model.Song) (*model.Song,error) {
	return s.songRepo.InsertSong(data)
}
