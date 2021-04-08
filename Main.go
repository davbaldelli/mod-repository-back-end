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

	jsonFile, _ := os.Open("credentials.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cred Credentials

	json.Unmarshal(byteValue, &cred)

	dsn := "host=192.168.0.113 user=" + cred.Username + " password=" + cred.Password + " dbname=mod-repo-2 port=5432 sslmode=disable TimeZone=Europe/Rome"
	dbase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	carRepo := postgresrepo.CarRepositoryImpl{Db: dbase}
	trackRepo := postgresrepo.TrackRepositoryImpl{Db: dbase}
	nationRepo := postgresrepo.NationsRepositoryImpl{Db: dbase}
	brandRepo := postgresrepo.BrandRepositoryImpl{Db: dbase}

	/*
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.113:27017").SetAuth(options.Credential{
		Username: "",
		Password: "",
	}))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	 := client.Database("mod_repo").Collection("tracks")
	carsCollection := client.Database("mod_repo").Collection("cars")
	 */
	web := routes.Web{
		CarHandler:    handlers.CarsHandlerImpl{CarCtrl: controllers.CarControllerImpl{Repo: carRepo}},
		TracksHandler: handlers.TrackHandlerImpl{TrackCtrl: controllers.TrackControllerImpl{Repo: trackRepo}},
		NationHandler: handlers.NationsHandlerImpl{CtrlNations: controllers.NationControllerImpl{Repo: nationRepo}},
		BrandsHandler: handlers.BrandsHandlerImpl{BrandCtrl: controllers.BrandControllerImpl{Repo: brandRepo}},
	}
	web.Listen("6316")
}
