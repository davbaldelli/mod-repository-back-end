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
	CarHandler     handlers.CarsHandler
	TracksHandler  handlers.TracksHandler
	NationHandler  handlers.NationsHandler
	BrandsHandler  handlers.BrandsHandler
	UsersHandler   handlers.UsersHandler
	AuthorsHandler handlers.AuthorsHandler
	LogsHandler handlers.LogsHandler
	Middleware handlers.Middleware
}

func (w Web) Listen() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/car/new", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.CarHandler.POSTNewCar, []string{"admin"}))).Methods("POST")
	router.HandleFunc("/car/update", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.CarHandler.UPDATECar, []string{"admin"}))).Methods("POST")
	router.HandleFunc("/car/all", w.Middleware.IsAuthorized(w.CarHandler.GETAllCars)).Methods("GET")
	router.HandleFunc("/car/type/all", w.Middleware.IsAuthorized(w.CarHandler.GETAllCarCategories)).Methods("GET")

	router.HandleFunc("/track/new", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.TracksHandler.POSTNewTrack, []string{"admin"}))).Methods("POST")
	router.HandleFunc("/track/update", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.TracksHandler.UPDATETrack, []string{"admin"}))).Methods("POST")
	router.HandleFunc("/track/all", w.Middleware.IsAuthorized(w.TracksHandler.GETAllTracks)).Methods("GET")

	router.HandleFunc("/log/car/all", w.Middleware.IsAuthorized(w.LogsHandler.GETAllCarLogs)).Methods("GET")
	router.HandleFunc("/log/track/all", w.Middleware.IsAuthorized(w.LogsHandler.GETAllTrackLogs)).Methods("GET")

	router.HandleFunc("/nation/brand/all", w.Middleware.IsAuthorized(w.NationHandler.GETAllBrandsNations)).Methods("GET")
	router.HandleFunc("/nation/track/all", w.Middleware.IsAuthorized(w.NationHandler.GETAllTracksNations)).Methods("GET")

	router.HandleFunc("/brand/all", w.Middleware.IsAuthorized(w.BrandsHandler.GETAllBrands)).Methods("GET")

	router.HandleFunc("/author/all", w.Middleware.IsAuthorized(w.AuthorsHandler.GETAllAuthors)).Methods("GET")
	router.HandleFunc("/car/author/all", w.Middleware.IsAuthorized(w.AuthorsHandler.GETCarAuthors)).Methods("GET")
	router.HandleFunc("/track/author/all", w.Middleware.IsAuthorized(w.AuthorsHandler.GETTrackAuthors)).Methods("GET")

	router.HandleFunc("/login", w.UsersHandler.LogIn).Methods("POST")
	router.HandleFunc("/signin", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.UsersHandler.SignIn, []string{"admin"}))).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://www.acmodrepository.com", "http://localhost:8080", "http://localhost:3000", "https://mods.davidebaldelli.it", "128.116.134.232"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("api.mod.davidebaldelli.it","api.acmodrepository.com"),
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
