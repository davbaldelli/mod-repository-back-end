package routes

import (
	"github.com/davide/ModRepository/routes/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	router.HandleFunc("/car/new", w.CarHandler.POSTNewCar).Methods("POST")
	router.HandleFunc("/car/all", w.CarHandler.GETAllCars).Methods("GET")
	router.HandleFunc("/car/type/all", w.CarHandler.GETAllCarCategories).Methods("GET")

	router.HandleFunc("/track/new", w.TracksHandler.POSTNewTrack).Methods("POST")
	router.HandleFunc("/track/all", w.TracksHandler.GETAllTracks).Methods("GET")

	router.HandleFunc("/nation/brand/all", w.NationHandler.GETAllBrandsNations).Methods("GET")
	router.HandleFunc("/nation/track/all", w.NationHandler.GETAllTracksNations).Methods("GET")

	router.HandleFunc("/brand/all", w.BrandsHandler.GETAllBrands).Methods("GET")


	router.HandleFunc("/author/all", w.AuthorsHandler.GETAllAuthors).Methods("GET")
	router.HandleFunc("/car/author/all", w.AuthorsHandler.GETCarAuthors).Methods("GET")
	router.HandleFunc("/track/author/all", w.AuthorsHandler.GETTrackAuthors).Methods("GET")

	router.HandleFunc("/login", w.UsersHandler.LogIn).Methods("POST")
	router.HandleFunc("/signin", w.UsersHandler.SignIn).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)
/*
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
*/
	log.Fatal(http.ListenAndServe(":6316", handler))

}



