package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"artistSpotify/artist"
	"artistSpotify/model"

	"github.com/gorilla/mux"
)

type ArtistHandler struct {
	artistUsecase artist.ArtistUsecase
}

func CreateArtistHandler(r *mux.Router, artistUsecase artist.ArtistUsecase) {
	artistHandler := ArtistHandler{artistUsecase}

	viewArtist := r.PathPrefix("/artist/genre").Subrouter().StrictSlash(true)
	viewArtist.HandleFunc("/{id}", artistHandler.findAllArtist).Methods(http.MethodGet)

	insertArtist := r.PathPrefix("/artist").Subrouter().StrictSlash(true)
	insertArtist.HandleFunc("/", artistHandler.insertArtist).Methods(http.MethodPost)
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

func (a *ArtistHandler) findAllArtist(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var getID = mux.Vars(req)
	id, err := strconv.Atoi(getID["id"])
	if err != nil {
		handlerError(w, err.Error())
		return
	}

	artist, err := a.artistUsecase.FindAllArtist(id)
	if err != nil {
		handlerError(w, err.Error())
		return
	}
	handlerSuccess(w, artist)
}

func (a *ArtistHandler) insertArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data model.Artist
	var body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		handlerError(w, "error failed when trying read body request")
		return
	}

	err = json.Unmarshal(body, &data)
	if err!=nil {
		handlerError(w, "Opps.. sorry something when wrong.")
		return
	}

	output, err := a.artistUsecase.InsertArtist(&data)
	if err !=nil {
		handlerError(w, err.Error())
		return
	}
	handlerSuccess(w, output)
}