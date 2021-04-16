package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"github.com/gorilla/mux"
	"net/http"
)

type TrackHandlerImpl struct {
	TrackCtrl controllers.TrackController
}

type getTracksByParam func(string) ([]entities.Track, error)

func (t TrackHandlerImpl) GETAllTracks(writer http.ResponseWriter, request *http.Request) {
	if tracks, err := t.TrackCtrl.GetAllTracks(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, tracks)
	}

}

func (t TrackHandlerImpl) GETTracksByNation(writer http.ResponseWriter, request *http.Request) {
	t.getTrackByParamResponse("nation", func(s string) ([]entities.Track, error) { return t.TrackCtrl.GetTracksByNation(s) }, writer, request)
}

func (t TrackHandlerImpl) GETTracksByLayoutType(writer http.ResponseWriter, request *http.Request) {
	t.getTrackByParamResponse("layoutType", func(s string) ([]entities.Track, error) { return t.TrackCtrl.GetTracksByLayoutType(s) }, writer, request)
}

func (t TrackHandlerImpl) GETTrackByName(writer http.ResponseWriter, request *http.Request) {
	t.getTrackByParamResponse("name", func(s string) ([]entities.Track, error) { return t.TrackCtrl.GetTracksByName(s) }, writer, request)
}

func (t TrackHandlerImpl) POSTNewTrack(writer http.ResponseWriter, request *http.Request) {
	track := entities.Track{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&track); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := t.TrackCtrl.AddTrack(track); err != nil {
		respondError(writer, http.StatusInternalServerError, fmt.Errorf("cannot insert new entity: %v ", err))
		return
	}

	respondJSON(writer, http.StatusCreated, track)
}

func (t TrackHandlerImpl) getTrackByParamResponse(paramString string, getTracks getTracksByParam, writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params[paramString]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param '"+paramString+"'"))
		return
	}

	if tracks, err := getTracks(param); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, tracks)
	}

}
