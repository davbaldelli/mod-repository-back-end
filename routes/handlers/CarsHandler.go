package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"github.com/gorilla/mux"
	"net/http"
)

type CarsHandlerImpl struct {
	CarCtrl controllers.CarController
}

func (c CarsHandlerImpl) GETCarByModel(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params["model"]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param model"))
		return
	}

	if car, err := c.CarCtrl.GetCarByModel(param); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, car)
	}
}

func (c CarsHandlerImpl) GETAllCarCategories(writer http.ResponseWriter, request *http.Request) {
	if categories, err := c.CarCtrl.GetAllCarCategories(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, categories)
	}
}

type getCarsByParam func(string) ([]entities.Car, error)

func (c CarsHandlerImpl) GETAllCars(writer http.ResponseWriter, request *http.Request) {
	if cars, err := c.CarCtrl.GetAllCars(request.Header.Get("Role") != string(entities.Base)); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, cars)
	}
}

func (c CarsHandlerImpl) GETCarsByNation(writer http.ResponseWriter, request *http.Request) {
	c.getCarsByParamResponse("nation", func(nation string) ([]entities.Car,error) { return c.CarCtrl.GetCarsByNation(nation, request.Header.Get("Role") != string(entities.Base)) }, writer, request)
}

func (c CarsHandlerImpl) GETCarsByModel(writer http.ResponseWriter, request *http.Request) {
	c.getCarsByParamResponse("model", func(model string) ([]entities.Car,error)  { return c.CarCtrl.GetCarsByModel(model, request.Header.Get("Role") != string(entities.Base)) }, writer, request)
}

func (c CarsHandlerImpl) GETCarsByBrand(writer http.ResponseWriter, request *http.Request) {
	c.getCarsByParamResponse("brand", func(brand string) ([]entities.Car,error)  { return c.CarCtrl.GetCarsByBrand(brand, request.Header.Get("Role") != string(entities.Base)) }, writer, request)
}

func (c CarsHandlerImpl) GETCarsByType(writer http.ResponseWriter, request *http.Request) {
	c.getCarsByParamResponse("category", func(category string) ([]entities.Car,error) { return c.CarCtrl.GetCarsByType(category, request.Header.Get("Role") != string(entities.Base)) }, writer, request)
}

func (c CarsHandlerImpl) POSTNewCar(writer http.ResponseWriter, request *http.Request) {
	car := entities.Car{}

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&car); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.CarCtrl.AddCar(car); err != nil {
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

	if cars, err := getCars(param) ; err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, cars)
	}
}
