package handlers

import (
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/presentation"
	"net/http"
)

type NationsHandlerImpl struct {
	CtrlNations controllers.NationController
}

func (n NationsHandlerImpl) GETAllTracksNations(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, http.StatusOK, presentation.OfAllNations(n.CtrlNations.GetAllTracksNations()))
}

func (n NationsHandlerImpl) GETAllBrandsNations(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, http.StatusOK, presentation.OfAllNations(n.CtrlNations.GetAllBrandsNations()))
}
