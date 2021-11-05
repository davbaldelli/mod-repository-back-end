package routes

import (
	"crypto/tls"
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/routes/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

type Route interface {
	Listen()
}

type Web struct {
	CarHandler    handlers.CarsHandler
	TracksHandler handlers.TracksHandler
	NationHandler handlers.NationsHandler
	BrandsHandler handlers.BrandsHandler
	UsersHandler  handlers.UsersHandler
	AuthorsHandler handlers.AuthorsHandler
}

func (w Web) Listen() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/car/new", handlers.IsAuthorized(handlers.IsAllowed(w.CarHandler.POSTNewCar, []string{string(entities.Admin)}))).Methods("POST")
	router.HandleFunc("/car/all", handlers.IsAuthorized(w.CarHandler.GETAllCars)).Methods("GET")
	router.HandleFunc("/car/nation/{nation}", handlers.IsAuthorized(w.CarHandler.GETCarsByNation)).Methods("GET")
	router.HandleFunc("/car/find/model/{model}", handlers.IsAuthorized(w.CarHandler.GETCarsByModel)).Methods("GET")
	router.HandleFunc("/car/brand/{brand}", handlers.IsAuthorized(w.CarHandler.GETCarsByBrand)).Methods("GET")
	router.HandleFunc("/car/category/{category}", handlers.IsAuthorized(w.CarHandler.GETCarsByType)).Methods("GET")
	router.HandleFunc("/car/type/all", handlers.IsAuthorized(w.CarHandler.GETAllCarCategories)).Methods("GET")
	router.HandleFunc("/car/model/{model}", handlers.IsAuthorized(w.CarHandler.GETCarByModel)).Methods("GET")

	router.HandleFunc("/track/new", handlers.IsAuthorized(handlers.IsAllowed(w.TracksHandler.POSTNewTrack, []string{string(entities.Admin)}))).Methods("POST")
	router.HandleFunc("/track/all", handlers.IsAuthorized(w.TracksHandler.GETAllTracks)).Methods("GET")
	router.HandleFunc("/track/nation/{nation}", handlers.IsAuthorized(w.TracksHandler.GETTracksByNation)).Methods("GET")
	router.HandleFunc("/track/layout/type/{layoutType}", handlers.IsAuthorized(w.TracksHandler.GETTracksByLayoutType)).Methods("GET")
	router.HandleFunc("/track/find/name/{name}", handlers.IsAuthorized(w.TracksHandler.GETTracksByName)).Methods("GET")
	router.HandleFunc("/track/tag/{tag}", handlers.IsAuthorized(w.TracksHandler.GETTracksByTag)).Methods("GET")
	router.HandleFunc("/track/name/{name}", handlers.IsAuthorized(w.TracksHandler.GETTrackByName)).Methods("GET")

	router.HandleFunc("/nation/brand/all", w.NationHandler.GETAllBrandsNations).Methods("GET")
	router.HandleFunc("/nation/track/all", w.NationHandler.GETAllTracksNations).Methods("GET")

	router.HandleFunc("/brand/all", handlers.IsAuthorized(w.BrandsHandler.GETAllBrands)).Methods("GET")
	router.HandleFunc("/brand/all/grouped/nation", handlers.IsAuthorized(w.BrandsHandler.GETAllBrandsGroupedByNation)).Methods("GET")
	router.HandleFunc("/brand/name/{name}", handlers.IsAuthorized(w.BrandsHandler.GETBrandByName)).Methods("GET")
	router.HandleFunc("/brand/nation/{nation}", handlers.IsAuthorized(w.BrandsHandler.GETBrandByNation)).Methods("GET")

	router.HandleFunc("/author/all", handlers.IsAuthorized(w.AuthorsHandler.GETAllAuthors)).Methods("GET")
	router.HandleFunc("/car/author/all", handlers.IsAuthorized(w.AuthorsHandler.GETCarAuthors)).Methods("GET")
	router.HandleFunc("/track/author/all", handlers.IsAuthorized(w.AuthorsHandler.GETTrackAuthors)).Methods("GET")

	router.HandleFunc("/login", w.UsersHandler.SignIn).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("api.mod.davidebaldelli.it"),
		Cache:      autocert.DirCache("certs"),
	}

	// create the server itself
	server := &http.Server{
		Addr:    ":https",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	log.Printf("Serving http/https for domains: api.mod.davidebaldelli.it")
	go func() {
		// serve HTTP, which will redirect automatically to HTTPS
		h := certManager.HTTPHandler(nil)
		log.Fatal(http.ListenAndServe(":http", h))
	}()

	log.Fatal(server.ListenAndServeTLS("", ""))

	//log.Fatal(http.ListenAndServe(":6316", handler))

}



