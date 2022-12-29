package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"net/http"
	"strconv"
)

type SkinsHandlerImpl struct {
	Ctrl controllers.SkinController
}

func (s SkinsHandlerImpl) GETCarSkins(w http.ResponseWriter, r *http.Request) {

	carId, err := strconv.ParseUint(r.URL.Query().Get("carId"), 10, 64)

	if err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("Missing 'carId' param. %v ", err))
		return
	}

	if skins, err := s.Ctrl.SelectCarSkins(uint(carId)); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("Error proccessing request: %v ", err))
		return
	} else {
		respondJSON(w, http.StatusOK, skins)
	}
}

func (s SkinsHandlerImpl) ADDSkin(w http.ResponseWriter, r *http.Request) {
	skin := entities.Skin{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&skin); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := s.Ctrl.AddSkin(skin); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error adding the new entity: %v ", err))
		return
	}

	respondJSON(w, http.StatusCreated, skin)
}

func (s SkinsHandlerImpl) UPDATESkin(w http.ResponseWriter, r *http.Request) {
	skin := entities.Skin{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&skin); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := s.Ctrl.UpdateSkin(skin); err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error updating the entity: %v ", err))
		return
	}

	respondJSON(w, http.StatusCreated, skin)
}
