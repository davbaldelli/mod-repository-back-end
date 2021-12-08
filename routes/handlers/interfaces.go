package handlers

import (
	"net/http"
)

type CarsHandler interface {
	GETAllCars(http.ResponseWriter, *http.Request)
	GETAllCarCategories(http.ResponseWriter, *http.Request)
	POSTNewCar(http.ResponseWriter, *http.Request)
	UPDATECar(http.ResponseWriter, *http.Request)
}

type TracksHandler interface {
	GETAllTracks(http.ResponseWriter, *http.Request)
	POSTNewTrack(http.ResponseWriter, *http.Request)
	UPDATETrack(http.ResponseWriter, *http.Request)
}

type LogsHandler interface {
	GETAllTrackLogs(http.ResponseWriter, *http.Request)
	GETAllCarLogs(http.ResponseWriter, *http.Request)
}

type NationsHandler interface {
	GETAllTracksNations(http.ResponseWriter, *http.Request)
	GETAllBrandsNations(http.ResponseWriter, *http.Request)
}

type BrandsHandler interface {
	GETAllBrands(http.ResponseWriter, *http.Request)
}

type UsersHandler interface {
	LogIn(http.ResponseWriter, *http.Request)
	SignIn(http.ResponseWriter, *http.Request)
}

type AuthorsHandler interface {
	GETAllAuthors(http.ResponseWriter, *http.Request)
	GETTrackAuthors(http.ResponseWriter, *http.Request)
	GETCarAuthors(http.ResponseWriter, *http.Request)
}

type Middleware interface {
	IsAuthorized(next http.HandlerFunc) http.HandlerFunc
	IsAllowed(next http.HandlerFunc, allowedRoles []string) http.HandlerFunc
}

type FirebaseHandler interface {
	SubscribeToTopic(http.ResponseWriter, *http.Request)
}