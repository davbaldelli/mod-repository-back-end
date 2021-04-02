package handlers

import (
	"github.com/davide/ModRepository/controllers"
	"net/http"
)

type NationsHandlerImpl struct {
	ctrlNations controllers.NationController
}

func (n NationsHandlerImpl) GETAllTracksNations(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, http.StatusOK, n.ctrlNations.GetAllTracksNations())
}

func (n NationsHandlerImpl) GETAllBrandsNations(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, http.StatusOK, n.ctrlNations.GetAllBrandsNations())
}
