package routes

import (
	"github.com/davide/ModRepository/routes/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Route interface {
	Listen(string)
}

type Web struct {
	CarHandler handlers.CarsHandler
}

func (w Web) Listen(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/car/new", w.CarHandler.POSTNewCar).Methods("POST")
	router.HandleFunc("/car/all", w.CarHandler.GETAllCars).Methods("GET")
	router.HandleFunc("/car/nation/{nation}", w.CarHandler.GETCarsByNation).Methods("GET")
	router.HandleFunc("/car/model/{model}", w.CarHandler.GETCarByModel).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
