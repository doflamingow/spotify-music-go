package usecase

import (
	"artistSpotify/genre"
	"artistSpotify/model"
)

type GenreUsecaseImpl struct {
	genreRepo genre.GenreRepo
}

func CreateGenreUsecaseImpl(genreRepo genre.GenreRepo) genre.GenreRepo {
	return &GenreUsecaseImpl{genreRepo}
}

func (call *GenreUsecaseImpl)GetAllGenre() (*[]model.Genre, error){
	return call.genreRepo.GetAllGenre()
}