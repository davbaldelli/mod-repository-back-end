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
	CarCtrl      controllers.CarController
	FirebaseCtrl controllers.FirebaseController
	DiscordBotCtrl controllers.DiscordBotController
}

func (c CarsHandlerImpl) GETAllCarCategories(writer http.ResponseWriter, _ *http.Request) {
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

func (c CarsHandlerImpl) POSTNewCar(writer http.ResponseWriter, request *http.Request) {
	car := entities.Car{}

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&car); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := c.CarCtrl.AddCar(&car); err != nil {
		respondError(writer, http.StatusInternalServerError, fmt.Errorf("cannot insert new entity: %v ", err))
		return
	}

	c.FirebaseCtrl.NotifyCarAdded(car)
	c.DiscordBotCtrl.NotifyCarAdded(car)

	respondJSON(writer, http.StatusCreated, car)
}

func (c CarsHandlerImpl) UPDATECar(writer http.ResponseWriter, request *http.Request) {
	car := entities.Car{}

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&car); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if versionChange, err := c.CarCtrl.UpdateCar(car); err != nil {
		respondError(writer, http.StatusInternalServerError, fmt.Errorf("cannot insert new entity: %v ", err))
		return
	} else if versionChange {
		c.FirebaseCtrl.NotifyCarUpdated(car)
		c.DiscordBotCtrl.NotifyCarUpdated(car)
	}


	respondJSON(writer, http.StatusOK, car)
}

func (c CarsHandlerImpl) getCarsByParamResponse(paramName string, getCars getCarsByParam, writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params[paramName]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param '"+paramName+"'"))
		return
	}

	if cars, err := getCars(param); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, cars)
	}
}
