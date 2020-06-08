package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"artistSpotify/artist"
	"artistSpotify/model"
)

type ArtistRepoImpl struct {
	db *sql.DB
}

func CreateArtistRepoMysqlImpl(db *sql.DB) artist.ArtistRepo {
	return &ArtistRepoImpl{db}
}

func (a *ArtistRepoImpl) FindAllArtist(id int) (*[]model.Artist, error) {
	query := "select artist.id,artist.name,artist.debut,artist.category,genre.name,artist.image_url from artist JOIN genre ON artist.genre_id = genre.id where genre_id = ?"

	rows, err := a.db.Query(query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[ArtistRepoImpl.FindAllArtist] Error when FindAllArtist artist : %w", err)
	}

	defer rows.Close()

	var result []model.Artist
	for rows.Next() {
		dataArtist := model.Artist{}
		err = rows.Scan(&dataArtist.ID, &dataArtist.Name, &dataArtist.Debut, &dataArtist.Category, &dataArtist.Genre, &dataArtist.ImageUrl)
		if err != nil {
			return nil, fmt.Errorf("[ArtistRepoImpl.FindAll] Error when scan rows artists : %w", err)
		}
		result = append(result, dataArtist)
	}

	return &result, nil
}

func (a * ArtistRepoImpl) InsertArtist(data *model.Artist) (*model.Artist, error){
	var query = "insert into artist (name, debut, genre_id, category, image_url) values (?,?,?,?,?)"

	result, err := a.db.Exec(query, data.Name,data.Debut, data.ID,  data.Category, data.ImageUrl)
	if err!=nil {
		return nil, fmt.Errorf("[ArtistRepoImpl.InsertArtist] Error when trying inserting data : %w ",err)
	}

	getID, err := result.LastInsertId()
	if err !=nil {
		return nil, fmt.Errorf("[ArtistRepoImpl] Error when trying to get last insert id : %w ", err)
	}
	output := model.Artist{ID: int(getID), Name: data.Name,Debut:data.Debut, Category: data.Category, Genre:data.Genre, ImageUrl:data.ImageUrl}
	return &output, nil
}
