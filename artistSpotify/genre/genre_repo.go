package genre

import "artistSpotify/model"

type GenreRepo interface{
	GetAllGenre() (*[]model.Genre, error)
}