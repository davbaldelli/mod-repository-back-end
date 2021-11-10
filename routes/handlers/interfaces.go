package handlers

import (
	"net/http"
)

type CarsHandler interface {
	GETAllCars(http.ResponseWriter, *http.Request)
	GETAllCarCategories(http.ResponseWriter, *http.Request)
	POSTNewCar(http.ResponseWriter, *http.Request)
}

type TracksHandler interface {
	GETAllTracks(http.ResponseWriter, *http.Request)
	POSTNewTrack(http.ResponseWriter, *http.Request)
}

type NationsHandler interface {
	GETAllTracksNations(http.ResponseWriter, *http.Request)
	GETAllBrandsNations(http.ResponseWriter, *http.Request)
}

type BrandsHandler interface {
	GETAllBrands(http.ResponseWriter, *http.Request)
}

type UsersHandler interface {
	POSTLogin(http.ResponseWriter, *http.Request)
	SignIn(http.ResponseWriter, *http.Request)
}

type AuthorsHandler interface {
	GETAllAuthors(http.ResponseWriter, *http.Request)
	GETTrackAuthors(http.ResponseWriter, *http.Request)
	GETCarAuthors(http.ResponseWriter, *http.Request)


}
