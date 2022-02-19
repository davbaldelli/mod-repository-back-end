package handlers

import (
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
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
		respondJSON(writer, http.StatusOK, brands)
	}
}
