package artist

import "artistSpotify/model"

type ArtistUsecase interface {
	FindAllArtist(id int) (*[]model.Artist, error)
	InsertArtist(data *model.Artist) (*model.Artist, error)
}
