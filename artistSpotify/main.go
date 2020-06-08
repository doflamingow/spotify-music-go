package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	handlerArtist "artistSpotify/artist/handler"
	repoArtist "artistSpotify/artist/repo"
	usecaseArtist "artistSpotify/artist/usecase"

	handlerGenre "artistSpotify/genre/handler"
	repoGenre "artistSpotify/genre/repo"
	usecaseGenre "artistSpotify/genre/usecase"

	handlerSong "artistSpotify/song/handler"
	repoSong "artistSpotify/song/repo"
	usecaseSong "artistSpotify/song/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	port := "8000"
	conStr := "root:root@tcp(127.0.0.1)/spotifay"

	db, err := sql.Open("mysql", conStr)
	defer db.Close()

	artistRepo := repoArtist.CreateArtistRepoMysqlImpl(db)
	artistUsecase := usecaseArtist.CreateArtistUsecaseImpl(artistRepo)

	genreRepo := repoGenre.CreateGenreRepoMysqlImpl(db)
	genreUsecase := usecaseGenre.CreateGenreUsecaseImpl(genreRepo)

	songRepo := repoSong.CreateSongRepoMysqlImpl(db)
	songUsecase := usecaseSong.CreateSongUsecaseImpl(songRepo)


	r := mux.NewRouter().StrictSlash(true)

	handlerArtist.CreateArtistHandler(r, artistUsecase)
	handlerGenre.CreateArtistHandler(r, genreUsecase)
	handlerSong.CreateSongHandler(r, songUsecase)

	fmt.Println("Starting web server at http://localhost:8000")
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
