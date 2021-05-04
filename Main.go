package main

import (
	"encoding/json"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/repositories/postgresrepo"
	"github.com/davide/ModRepository/routes"
	"github.com/davide/ModRepository/routes/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

type Credentials struct {
	Username string
	Password string
}

func main() {

	jsonFile, err := os.Open("credentials.json")

	if err != nil {
		log.Fatal(err)
	}

	byteValue, err1 := ioutil.ReadAll(jsonFile)

	if err1 != nil {
		log.Fatal(err)
	}

	var cred Credentials

	if err2 := json.Unmarshal(byteValue, &cred); err2 != nil {
		log.Fatal(err)
	}

	dsn := "host=127.0.0.1 user=" + cred.Username + " password=" + cred.Password + " dbname=mod_repo port=5432 sslmode=disable TimeZone=Europe/Rome"
	dbase, err3 := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err3 != nil {
		log.Fatal(err)
	}

	carRepo := postgresrepo.CarRepositoryImpl{Db: dbase}
	trackRepo := postgresrepo.TrackRepositoryImpl{Db: dbase}
	nationRepo := postgresrepo.NationsRepositoryImpl{Db: dbase}
	brandRepo := postgresrepo.BrandRepositoryImpl{Db: dbase}
	userRepo := postgresrepo.UserRepositoryImpl{Db: dbase}
	authorRepo := postgresrepo.AuthorsRepositoryImpl{Db: dbase}

	web := routes.Web{
		CarHandler:    handlers.CarsHandlerImpl{CarCtrl: controllers.CarControllerImpl{Repo: carRepo}},
		TracksHandler: handlers.TrackHandlerImpl{TrackCtrl: controllers.TrackControllerImpl{Repo: trackRepo}},
		NationHandler: handlers.NationsHandlerImpl{CtrlNations: controllers.NationControllerImpl{Repo: nationRepo}},
		BrandsHandler: handlers.BrandsHandlerImpl{BrandCtrl: controllers.BrandControllerImpl{Repo: brandRepo}},
		UsersHandler:  handlers.UserHandlerImpl{UserCtrl: controllers.UserControllerImpl{Repo: userRepo}},
		AuthorsHandler: handlers.AuthorHandlerImpl{AuthorsCtrl: controllers.AuthorsControllerImpl{Repo: authorRepo}},
	}
	web.Listen()
}
