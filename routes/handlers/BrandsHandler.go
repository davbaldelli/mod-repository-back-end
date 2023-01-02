package handlers

import (
	"github.com/davide/ModRepository/controllers"
	"net/http"
)

type BrandsHandlerImpl struct {
	BrandCtrl controllers.BrandController
}

func (b BrandsHandlerImpl) GETAllBrands(writer http.ResponseWriter, _ *http.Request) {
	if brands, err := b.BrandCtrl.GetAllBrands(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, brands)
	}
}
