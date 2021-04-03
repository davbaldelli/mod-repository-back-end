package handlers

import (
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net/http"
)

type TrackHandlerImpl struct {
	TrackCtrl controllers.TrackController
}

type getTracksByParam func(string) []models.Track

func (t TrackHandlerImpl) GETAllTracks(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, http.StatusOK, t.TrackCtrl.GetAllTracks())
}

func (t TrackHandlerImpl) GETTracksByNation(writer http.ResponseWriter, request *http.Request) {
	t.getTrackByParamResponse("nation", func(s string) []models.Track { return t.TrackCtrl.GetTracksByNation(s) }, writer, request)
}

func (t TrackHandlerImpl) GETTracksByLayoutType(writer http.ResponseWriter, request *http.Request) {
	t.getTrackByParamResponse("layoutType", func(s string) []models.Track { return t.TrackCtrl.GetTracksByLayoutType(s) }, writer, request)
}

func (t TrackHandlerImpl) GETTrackByName(writer http.ResponseWriter, request *http.Request) {
	t.getTrackByParamResponse("name", func(s string) []models.Track { return t.TrackCtrl.GetTracksByName(s) }, writer, request)
}

func (t TrackHandlerImpl) POSTNewTrack(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error parsing request to post form: %v ", err))
		return
	}

	track := models.Track{}

	decoder := schema.NewDecoder()
	if err := decoder.Decode(&track, request.PostForm); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := t.TrackCtrl.AddTrack(track.Name, track.DownloadLink, track.Layouts, track.Location); err != nil {
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

	respondJSON(writer, http.StatusOK, getTracks(param))
}
