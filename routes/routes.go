package routes

import (
	"crypto/tls"
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
	//router.Use(handlers.IsAuthorized)
	router.HandleFunc("/car/new", w.CarHandler.POSTNewCar).Methods("POST")
	router.HandleFunc("/car/all", w.CarHandler.GETAllCars).Methods("GET")
	router.HandleFunc("/car/nation/{nation}", w.CarHandler.GETCarsByNation).Methods("GET")
	router.HandleFunc("/car/find/model/{model}", w.CarHandler.GETCarsByModel).Methods("GET")
	router.HandleFunc("/car/brand/{brand}", w.CarHandler.GETCarsByBrand).Methods("GET")
	router.HandleFunc("/car/category/{category}", w.CarHandler.GETCarsByType).Methods("GET")
	router.HandleFunc("/car/type/all", w.CarHandler.GETAllCarCategories).Methods("GET")
	router.HandleFunc("/car/model/{model}", w.CarHandler.GETCarByModel).Methods("GET")

	router.HandleFunc("/track/new", w.TracksHandler.POSTNewTrack).Methods("POST")
	router.HandleFunc("/track/all", w.TracksHandler.GETAllTracks).Methods("GET")
	router.HandleFunc("/track/nation/{nation}", w.TracksHandler.GETTracksByNation).Methods("GET")
	router.HandleFunc("/track/layout/type/{layoutType}", w.TracksHandler.GETTracksByLayoutType).Methods("GET")
	router.HandleFunc("/track/find/name/{name}", w.TracksHandler.GETTracksByName).Methods("GET")
	router.HandleFunc("/track/tag/{tag}", w.TracksHandler.GETTracksByTag).Methods("GET")
	router.HandleFunc("/track/name/{name}", w.TracksHandler.GETTrackByName).Methods("GET")

	router.HandleFunc("/nation/brand/all", w.NationHandler.GETAllBrandsNations).Methods("GET")
	router.HandleFunc("/nation/track/all", w.NationHandler.GETAllTracksNations).Methods("GET")

	router.HandleFunc("/brand/all", w.BrandsHandler.GETAllBrands).Methods("GET")
	router.HandleFunc("/brand/all/grouped/nation", w.BrandsHandler.GETAllBrandsGroupedByNation).Methods("GET")
	router.HandleFunc("/brand/name/{name}", w.BrandsHandler.GETBrandByName).Methods("GET")
	router.HandleFunc("/brand/nation/{nation}", w.BrandsHandler.GETBrandByNation).Methods("GET")

	router.HandleFunc("/author/all", w.AuthorsHandler.GETAllAuthors).Methods("GET")
	router.HandleFunc("/car/author/all", w.AuthorsHandler.GETCarAuthors).Methods("GET")
	router.HandleFunc("/track/author/all", w.AuthorsHandler.GETTrackAuthors).Methods("GET")

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



