package handlers

import (
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/models/presentation"
	"github.com/gorilla/mux"
	"net/http"
)

type getBrandsByParam func(string) []entities.CarBrand

type BrandsHandlerImpl struct {
	BrandCtrl controllers.BrandController
}

func (b BrandsHandlerImpl) GETAllBrands(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, http.StatusOK, presentation.OfAllBrands(b.BrandCtrl.GetAllBrands()))
}

func (b BrandsHandlerImpl) GETBrandByNation(writer http.ResponseWriter, request *http.Request) {
	b.getBrandByParamResponse("nation", func(s string) []entities.CarBrand { return b.BrandCtrl.GetBrandByNation(s) }, writer, request)
}

func (b BrandsHandlerImpl) GETBrandByName(writer http.ResponseWriter, request *http.Request) {
	b.getBrandByParamResponse("name", func(s string) []entities.CarBrand { return b.BrandCtrl.GetBrandByName(s) }, writer, request)
}


func (b BrandsHandlerImpl) getBrandByParamResponse(paramString string, getBrands getBrandsByParam, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params[paramString]

	if param == "" {
		respondError(w, http.StatusBadRequest, fmt.Errorf("missing param '"+paramString+"'"))
		return
	}

	respondJSON(w, http.StatusOK, presentation.OfAllBrands(getBrands(param)))
}
