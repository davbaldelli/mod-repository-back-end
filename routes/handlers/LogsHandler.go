package handlers

import (
	"github.com/davide/ModRepository/controllers"
	"net/http"
)

type LogsHandlerImpl struct {
	Ctrl controllers.LogController
}

func (l LogsHandlerImpl) GETAllTrackLogs(writer http.ResponseWriter, request *http.Request) {
	if logs, err := l.Ctrl.GetTrackLogs(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, logs)
	}
}

func (l LogsHandlerImpl) GETAllCarLogs(writer http.ResponseWriter, request *http.Request) {
	if logs, err := l.Ctrl.GetCarLogs(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, logs)
	}
}
