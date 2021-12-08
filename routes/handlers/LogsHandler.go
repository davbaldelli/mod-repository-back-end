package handlers

import (
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"net/http"
)

type LogsHandlerImpl struct {
	Ctrl controllers.LogController
}

func (l LogsHandlerImpl) GETAllTrackLogs(writer http.ResponseWriter, request *http.Request) {
	if logs, err := l.Ctrl.GetTrackLogs(request.Header.Get("Role") != string(entities.Base)); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, logs)
	}
}

func (l LogsHandlerImpl) GETAllCarLogs(writer http.ResponseWriter, request *http.Request) {
	if logs, err := l.Ctrl.GetCarLogs(request.Header.Get("Role") != string(entities.Base)); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, logs)
	}
}
