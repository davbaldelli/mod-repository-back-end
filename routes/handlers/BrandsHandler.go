package handlers

import (
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/models/presentation"
	"github.com/gorilla/mux"
	"net/http"
)

type getBrandsByParam func(string) ([]entities.CarBrand, error)

type BrandsHandlerImpl struct {
	BrandCtrl controllers.BrandController
}

func (b BrandsHandlerImpl) GETAllBrands(writer http.ResponseWriter, request *http.Request) {
	if brands, err := b.BrandCtrl.GetAllBrands(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, presentation.OfAllBrands(brands))
	}
}

func (b BrandsHandlerImpl) GETBrandByNation(writer http.ResponseWriter, request *http.Request) {
	b.getBrandByParamResponse("nation", func(s string) ([]entities.CarBrand, error) { return b.BrandCtrl.GetBrandsByNation(s) }, writer, request)
}

func (b BrandsHandlerImpl) GETBrandByName(writer http.ResponseWriter, request *http.Request) {
	b.getBrandByParamResponse("name", func(s string) ([]entities.CarBrand, error) { return b.BrandCtrl.GetBrandsByName(s) }, writer, request)
}


func (b BrandsHandlerImpl) getBrandByParamResponse(paramString string, getBrands getBrandsByParam, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params[paramString]

	if param == "" {
		respondError(w, http.StatusBadRequest, fmt.Errorf("missing param '"+paramString+"'"))
		return
	}

	if brands, err := getBrands(param); err!=nil{
		respondError(w, http.StatusInternalServerError, err)
	} else {
		respondJSON(w, http.StatusOK, presentation.OfAllBrands(brands))
	}
}
