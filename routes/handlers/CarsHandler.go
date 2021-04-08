package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/models/presentation"
	"github.com/gorilla/mux"
	"net/http"
)

type CarsHandlerImpl struct {
	CarCtrl controllers.CarController
}

type getCarsByParam func(string) []entities.Car

func (c CarsHandlerImpl) GETAllCars(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, http.StatusOK, presentation.OfAllCars(c.CarCtrl.GetAllCars()))
}

func (c CarsHandlerImpl) GETCarsByNation(writer http.ResponseWriter, request *http.Request) {
	c.getCarsByParamResponse("nation", func(s string) []entities.Car { return c.CarCtrl.GetCarsByNation(s) }, writer, request)
}

func (c CarsHandlerImpl) GETCarByModel(writer http.ResponseWriter, request *http.Request) {
	c.getCarsByParamResponse("model", func(s string) []entities.Car { return c.CarCtrl.GetCarByModel(s) }, writer, request)
}

func (c CarsHandlerImpl) GETCarsByBrand(writer http.ResponseWriter, request *http.Request) {
	c.getCarsByParamResponse("brand", func(s string) []entities.Car { return c.CarCtrl.GetCarsByBrand(s) }, writer, request)
}

func (c CarsHandlerImpl) GETCarsByType(writer http.ResponseWriter, request *http.Request) {
	c.getCarsByParamResponse("category", func(s string) []entities.Car { return c.CarCtrl.GetCarsByType(s) }, writer, request)
}

func (c CarsHandlerImpl) POSTNewCar(writer http.ResponseWriter, request *http.Request) {
	car := entities.Car{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&car); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.CarCtrl.AddCar(car.ModelName, car.DownloadLink, car.Brand, car.Categories, car.Year); err != nil {
		respondError(writer, http.StatusInternalServerError, fmt.Errorf("cannot insert new entity: %v ", err))
		return
	}

	respondJSON(writer, http.StatusCreated, car)

}

func (c CarsHandlerImpl) getCarsByParamResponse(paramName string, getCars getCarsByParam, writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params[paramName]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param '"+paramName+"'"))
		return
	}

	respondJSON(writer, http.StatusOK, presentation.OfAllCars(getCars(param)))
}
