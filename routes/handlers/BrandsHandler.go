package handlers

import (
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net/http"
)

type getBrandsByParam func(string) []models.CarBrand

type BrandsHandlerImpl struct {
	BrandCtrl controllers.BrandController
}

func (b BrandsHandlerImpl) GETAllBrands(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, http.StatusOK, b.BrandCtrl.GetAllBrands())
}

func (b BrandsHandlerImpl) GETBrandByNation(writer http.ResponseWriter, request *http.Request) {
	b.getBrandByParamResponse("nation", func(s string) []models.CarBrand { return b.BrandCtrl.GetBrandByNation(s) }, writer, request)
}

func (b BrandsHandlerImpl) GETBrandByName(writer http.ResponseWriter, request *http.Request) {
	b.getBrandByParamResponse("name", func(s string) []models.CarBrand { return b.BrandCtrl.GetBrandByName(s) }, writer, request)
}

func (b BrandsHandlerImpl) POSTNewBrand(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error parsing request to post form: %v ", err))
		return
	}

	brand := models.CarBrand{}

	decoder := schema.NewDecoder()
	if err := decoder.Decode(&brand, request.PostForm); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := b.BrandCtrl.AddBrand(brand.Name, brand.Nation); err != nil {
		respondError(writer, http.StatusInternalServerError, fmt.Errorf("cannot insert new entity: %v ", err))
		return
	}

	respondJSON(writer, http.StatusCreated, brand)
}

func (b BrandsHandlerImpl) getBrandByParamResponse(paramString string, getBrands getBrandsByParam, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params[paramString]

	if param == "" {
		respondError(w, http.StatusBadRequest, fmt.Errorf("missing param '"+paramString+"'"))
		return
	}

	respondJSON(w, http.StatusOK, getBrands(param))
}
