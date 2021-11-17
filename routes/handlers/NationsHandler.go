package handlers

import (
	"github.com/davide/ModRepository/controllers"
	"net/http"
)

type NationsHandlerImpl struct {
	CtrlNations controllers.NationController
}

func (n NationsHandlerImpl) GETAllTracksNations(writer http.ResponseWriter, request *http.Request) {
	if nations, err := n.CtrlNations.GetAllTracksNations(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, nations)
	}
}

func (n NationsHandlerImpl) GETAllBrandsNations(writer http.ResponseWriter, request *http.Request) {
	if nations, err := n.CtrlNations.GetAllBrandsNations(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, nations)
	}
}
