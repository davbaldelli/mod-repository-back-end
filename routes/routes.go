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
	CarHandler    handlers.CarsHandler
	TracksHandler handlers.TracksHandler
}

func (w Web) Listen(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/car/new", w.CarHandler.POSTNewCar).Methods("POST")
	router.HandleFunc("/car/all", w.CarHandler.GETAllCars).Methods("GET")
	router.HandleFunc("/car/nation/{nation}", w.CarHandler.GETCarsByNation).Methods("GET")
	router.HandleFunc("/car/model/{model}", w.CarHandler.GETCarByModel).Methods("GET")
	router.HandleFunc("/car/brand/{brand}", w.CarHandler.GETCarsByBrand).Methods("GET")
	router.HandleFunc("/car/category/{category}", w.CarHandler.GETCarsByType).Methods("GET")

	router.HandleFunc("/track/new", w.TracksHandler.POSTNewTrack).Methods("POST")
	router.HandleFunc("/track/all", w.TracksHandler.GETAllTracks).Methods("GET")
	router.HandleFunc("/track/nation/{nation}", w.TracksHandler.GETTracksByNation).Methods("GET")
	router.HandleFunc("/track/layout/type/{layoutType}", w.TracksHandler.GETTracksByLayoutType).Methods("GET")
	router.HandleFunc("/track/name/{name}", w.TracksHandler.GETTrackByName).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
