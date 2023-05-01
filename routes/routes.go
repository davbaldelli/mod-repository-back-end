package routes

import (
	"crypto/tls"
	"github.com/davide/ModRepository/routes/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
	"sync"
)

type Route interface {
	Listen()
}

type Web struct {
	CarHandler      handlers.CarsHandler
	TracksHandler   handlers.TracksHandler
	NationHandler   handlers.NationsHandler
	BrandsHandler   handlers.BrandsHandler
	UsersHandler    handlers.UsersHandler
	AuthorsHandler  handlers.AuthorsHandler
	LogsHandler     handlers.LogsHandler
	ServersHandler  handlers.ServersHandler
	Middleware      handlers.Middleware
	FirebaseHandler handlers.FirebaseHandler
	SkinsHandler    handlers.SkinHandler
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

	router.HandleFunc("/skin", w.Middleware.IsAuthorized(w.SkinsHandler.GETCarSkins)).Methods("GET")
	router.HandleFunc("/skin/all", w.Middleware.IsAuthorized(w.SkinsHandler.GETAllSkins)).Methods("GET")
	router.HandleFunc("/skin/add", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.SkinsHandler.ADDSkin, []string{"admin"}))).Methods("POST")
	router.HandleFunc("/skin/update", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.SkinsHandler.UPDATESkin, []string{"admin"}))).Methods("POST")

	router.HandleFunc("/fsr/server1/all", w.Middleware.IsAuthorized(w.ServersHandler.GETAllServers)).Methods("GET")
	router.HandleFunc("/fsr/server1/update", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.ServersHandler.UPDATEServer, []string{"admin", "fsrteam"}))).Methods("POST")
	router.HandleFunc("/fsr/server1/add", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.ServersHandler.ADDServer, []string{"admin", "fsrteam"}))).Methods("POST")
	router.HandleFunc("/fsr/server1/delete", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.ServersHandler.DELETEServer, []string{"admin", "fsrteam"}))).Methods("POST")

	router.HandleFunc("/login", w.UsersHandler.LogIn).Methods("POST")
	router.HandleFunc("/signin", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.UsersHandler.SignIn, []string{"admin"}))).Methods("POST")
	router.HandleFunc("/user/updatepassword", w.Middleware.IsAuthorized(w.Middleware.IsAllowed(w.UsersHandler.UpdatePassword, []string{"admin"}))).Methods("POST")

	router.HandleFunc("/notification/register", w.Middleware.IsAuthorized(w.FirebaseHandler.SubscribeToTopic)).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://www.acmodrepository.com", "http://localhost:8080", "http://localhost:3000", "https://mods.davidebaldelli.it", "128.116.134.232", "https://fsr-dev--nuxt-acmodrepo.netlify.app"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("home.davidebaldelli.it", "api.acmodrepository.com"),
		Cache:      autocert.DirCache("certs"),
	}

	// create the server1 itself
	server1 := &http.Server{
		Addr:    ":443",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	server2 := &http.Server{
		Addr:    ":6316",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	var wg sync.WaitGroup

	wg.Add(3)

	log.Printf("Serving :6316 for domains: home.davidebaldelli.it , api.acmodrepository.com")

	go func() {
		defer wg.Done()
		log.Fatal(server1.ListenAndServeTLS("", ""))
	}()

	go func() {
		defer wg.Done()
		// serve HTTP, which will redirect automatically to HTTPS
		h := certManager.HTTPHandler(nil)
		log.Fatal(http.ListenAndServe(":http", h))
	}()

	go func() {
		defer wg.Done()
		log.Fatal(server2.ListenAndServeTLS("", ""))
	}()

	wg.Wait()

	//log.Fatal(http.ListenAndServe(":6316", handler))

}
