package repo

import (
	"artistSpotify/genre"
	"artistSpotify/model"
	"database/sql"
	"errors"
	"fmt"
)

type GenreRepoImpl struct {
	db *sql.DB
}

func CreateGenreRepoMysqlImpl(db *sql.DB) genre.GenreRepo {
	return &GenreRepoImpl{db}
}

func (g *GenreRepoImpl)GetAllGenre() (*[]model.Genre, error) {
	var query = "select id,name from genre"

	rows, err := g.db.Query(query)

	if err!=nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil,nil
		}
		return nil, fmt.Errorf("[GenreRepoImpl.GetAllGenre] Error when get all genre : %w", err)
	}
	defer rows.Close()

	var result []model.Genre
	for rows.Next() {
		dataGenre := model.Genre{}
		err := rows.Scan(&dataGenre.ID,&dataGenre.Name)
		if err!=nil {
			return nil, fmt.Errorf("[GenreRepoImpl.GetAllGenre] Error when scan all genre : %w", err)
		}
		result = append(result, dataGenre)
	}
	return &result, nil
}