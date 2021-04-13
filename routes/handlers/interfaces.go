package handlers

import (
	"net/http"
)

type CarsHandler interface {
	GETAllCars(http.ResponseWriter, *http.Request)
	GETCarsByNation(http.ResponseWriter, *http.Request)
	GETCarByModel(http.ResponseWriter, *http.Request)
	GETCarsByBrand(http.ResponseWriter, *http.Request)
	GETCarsByType(http.ResponseWriter, *http.Request)
	GETAllCarCategories(http.ResponseWriter, *http.Request)

	POSTNewCar(http.ResponseWriter, *http.Request)
}

type TracksHandler interface {
	GETAllTracks(http.ResponseWriter, *http.Request)
	GETTracksByNation(http.ResponseWriter, *http.Request)
	GETTracksByLayoutType(http.ResponseWriter, *http.Request)
	GETTrackByName(http.ResponseWriter, *http.Request)

	POSTNewTrack(http.ResponseWriter, *http.Request)
}

type NationsHandler interface {
	GETAllTracksNations(http.ResponseWriter, *http.Request)
	GETAllBrandsNations(http.ResponseWriter, *http.Request)
}

type BrandsHandler interface {
	GETAllBrands(http.ResponseWriter, *http.Request)
	GETAllBrandsGroupedByNation(http.ResponseWriter, *http.Request)
	GETBrandByNation(http.ResponseWriter, *http.Request)
	GETBrandByName(http.ResponseWriter, *http.Request)
}

type UsersHandler interface {
	POSTLogin(http.ResponseWriter, *http.Request)
}
