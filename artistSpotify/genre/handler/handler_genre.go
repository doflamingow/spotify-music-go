package handler

import (
	"artistSpotify/genre"
	"artistSpotify/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type GenreHandler struct {
	genreUsecase genre.GenreUsecase
}

func CreateArtistHandler(r *mux.Router, usecase genre.GenreUsecase) {
	genreHandler := GenreHandler{usecase}

	viewGenre := r.PathPrefix("/artist/genre").Subrouter().StrictSlash(true)
	viewGenre.HandleFunc("/", genreHandler.GetAllGenre).Methods(http.MethodGet)
}


func handlerSuccess(w http.ResponseWriter, data interface{}) {
	returnData := model.ResponseWrapper{
		Success: true,
		Message: "SUCCESS",
		Data:    data,
	}
	json.NewEncoder(w).Encode(returnData)
}

func handlerError(w http.ResponseWriter, message string) {
	data := model.ResponseWrapper{
		Success: true,
		Message: message,
	}
	json.NewEncoder(w).Encode(data)
}

func (call *GenreHandler) GetAllGenre(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	artist, err := call.genreUsecase.GetAllGenre()
	if err != nil {
		handlerError(w, err.Error())
		return
	}
	handlerSuccess(w, artist)
}
