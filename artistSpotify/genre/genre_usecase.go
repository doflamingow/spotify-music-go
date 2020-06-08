package genre

import "artistSpotify/model"

type GenreUsecase interface {
	GetAllGenre() (*[]model.Genre, error)
}