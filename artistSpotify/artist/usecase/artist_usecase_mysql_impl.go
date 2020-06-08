package usecase

import (
	"artistSpotify/artist"
	"artistSpotify/model"
)

type ArtistUsecaseImpl struct {
	artistRepo artist.ArtistRepo
}

func CreateArtistUsecaseImpl(artistRepo artist.ArtistRepo) artist.ArtistUsecase {
	return &ArtistUsecaseImpl{artistRepo}
}

func (a *ArtistUsecaseImpl) FindAllArtist(id int) (*[]model.Artist, error) {
	return a.artistRepo.FindAllArtist(id)
}

func (a *ArtistUsecaseImpl) InsertArtist(data *model.Artist) (*model.Artist, error)  {
	return a.artistRepo.InsertArtist(data)
}