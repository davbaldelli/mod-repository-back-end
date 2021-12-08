package handlers

import (
	"encoding/json"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"github.com/gorilla/mux"
	"net/http"
)

type TrackHandlerImpl struct {
	TrackCtrl    controllers.TrackController
	FirebaseCtrl controllers.FirebaseController
}

type getTracksByParam func(string) ([]entities.Track, error)

func (t TrackHandlerImpl) GETAllTracks(writer http.ResponseWriter, request *http.Request) {
	if tracks, err := t.TrackCtrl.GetAllTracks(request.Header.Get("Role") != string(entities.Base)); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, tracks)
	}

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

	message := &messaging.Message{
		Webpush: &messaging.WebpushConfig{Notification: &messaging.WebpushNotification{Actions: []*messaging.WebpushNotificationAction{
			{Action: "track_added", Title: "Check It Out!"},
		}}},
		Notification: &messaging.Notification{Title: fmt.Sprintf("%v has been added to repository", track.Name), Body: "A resource has been added", ImageURL: "https://imgur.com/0GuN24g"},
		Topic:        "modsUpdates",
	}
	t.FirebaseCtrl.Notify(message)

	respondJSON(writer, http.StatusCreated, track)
}

func (t TrackHandlerImpl) UPDATETrack(writer http.ResponseWriter, request *http.Request) {
	track := entities.Track{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&track); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if versionChange, err := t.TrackCtrl.UpdateTrack(track); err != nil {
		respondError(writer, http.StatusInternalServerError, fmt.Errorf("cannot insert new entity: %v ", err))
		return
	} else if versionChange {
		message := &messaging.Message{
			Webpush: &messaging.WebpushConfig{Notification: &messaging.WebpushNotification{Actions: []*messaging.WebpushNotificationAction{
				{Action: "track_updated", Title: "Check It Out!"},
			}}},
			Notification: &messaging.Notification{Title: fmt.Sprintf("%v has been updated", track.Name), Body: "A resource has been updated", ImageURL: "https://imgur.com/0GuN24g"},
			Topic:        "modsUpdates",
		}
		t.FirebaseCtrl.Notify(message)
	}

	respondJSON(writer, http.StatusOK, track)
}

func (t TrackHandlerImpl) getTrackByParamResponse(paramString string, getTracks getTracksByParam, writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params[paramString]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param '"+paramString+"'"))
		return
	}

	if tracks, err := getTracks(param); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, tracks)
	}

}
