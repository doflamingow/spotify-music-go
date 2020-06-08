package repo

import (
	"artistSpotify/model"
	"artistSpotify/song"
	"database/sql"
	"errors"
	"fmt"
)

type SongRepoImpl struct {
	db *sql.DB
}

func CreateSongRepoMysqlImpl (db *sql.DB) song.SongRepo {
	return &SongRepoImpl{db}
}

func (s *SongRepoImpl) GetAllSong(id int) (*[]model.Song, error) {
	var query = "select song.id, artist.name ,song.song_name ,song.imageUrl from song JOIN artist ON song.artist_id = artist.id where song.artist_id = ?"

	rows, err := s.db.Query(query, id)
	if err!=nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil,nil
		}

		return nil, fmt.Errorf("[SongRepoImpl.GetAllSong] Error when get all song : %w ", err)
	}
	defer rows.Close()

	var result []model.Song
	for rows.Next() {
		dataSong := model.Song{}
		err := rows.Scan(&dataSong.ID,&dataSong.Name,&dataSong.SongName,&dataSong.ImageUrl)
		if err!=nil {
			return nil, fmt.Errorf("[SongRepoImpl.GetAllSong] Error when scan song : %w ",err)
		}
		result = append(result, dataSong)
	}
	return &result, nil
}

func (s *SongRepoImpl) InsertSong(data model.Song) (*model.Song,error) {
	var query = "insert into song(artist_id ,song_name , imageUrl) values (? ,? , ?)"

	result ,err := s.db.Exec(query,data.ID,data.SongName, data.ImageUrl)
	if err!=nil {
		return nil, fmt.Errorf("[SongRepoImpl.InsertSong] Error when trying inserting a data : %w ", err)
	}

	getID, err := result.LastInsertId()
	if err!=nil {
		return nil, fmt.Errorf("[SongRepoImpl.InsertSong] Error when trying get lsatinser id : %w ",err)
	}
	output := model.Song{ID: int(getID), Name: data.Name,SongName: data.SongName, ImageUrl: data.ImageUrl}
	return &output,nil
}