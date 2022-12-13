package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"net/http"
)

type ServersHandlerImpl struct {
	Ctrl controllers.ServersController
}

func (s ServersHandlerImpl) ADDServer(w http.ResponseWriter, r *http.Request) {
	server := entities.Server{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&server); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := s.Ctrl.AddServer(server); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding the server: %v", err))
	} else {
		respondJSON(w, http.StatusOK, server)
	}
}

func (s ServersHandlerImpl) UPDATEServer(w http.ResponseWriter, r *http.Request) {
	server := entities.Server{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&server); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := s.Ctrl.UpdateServer(server); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error updating the server: %v", err))
	} else {
		respondJSON(w, http.StatusOK, server)
	}
}

func (s ServersHandlerImpl) DELETEServer(w http.ResponseWriter, r *http.Request) {
	server := entities.Server{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&server); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := s.Ctrl.DeleteServer(server); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error updating the server: %v", err))
	} else {
		respondJSON(w, http.StatusOK, "server deleted successfully")
	}
}

func (s ServersHandlerImpl) GETAllServers(writer http.ResponseWriter, _ *http.Request) {
	if cars, err := s.Ctrl.GetAllServers(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, cars)
	}
}
