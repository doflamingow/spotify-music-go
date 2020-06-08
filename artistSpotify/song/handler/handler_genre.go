package handler

import (
	"artistSpotify/model"
	"artistSpotify/song"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SongHandler struct {
	songUsercase song.SongUsecase
}

func CreateSongHandler(r *mux.Router, usecase song.SongUsecase) {
	songHandler := SongHandler{usecase}

	newRouterSong := r.PathPrefix("/artist/song").Subrouter().StrictSlash(true)

	newRouterSong.HandleFunc("/{id}", songHandler.getSong).Methods(http.MethodGet)
	newRouterSong.HandleFunc("/", songHandler.insertSong).Methods(http.MethodPost)

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

func (call *SongHandler) getSong(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var getID = mux.Vars(req)
	id, err := strconv.Atoi(getID["id"])
	if err != nil {
		handlerError(w, err.Error())
		return
	}
	artist, err := call.songUsercase.GetAllSong(id)
	if err != nil {
		handlerError(w, err.Error())
		return
	}
	handlerSuccess(w, artist)
}

func (call *SongHandler) insertSong(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data model.Song
	var body, err = ioutil.ReadAll(req.Body)
	if err!=nil {
		handlerError(w,"something wrong failed to read body request..")
		return
	}

	err = json.Unmarshal(body,&data)
	if err!=nil{
		handlerError(w,"Opps.. sorry something when wrong..."+err.Error())
		return
	}
	output, err := call.songUsercase.InsertSong(data)
	if err!=nil {
		handlerError(w, err.Error())
		return
	}
	handlerSuccess(w, output)
	return
}