package main

import (
	"context"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/repositories/mongorepo"
	"github.com/davide/ModRepository/repositories/postgresrepo"
	"github.com/davide/ModRepository/routes"
	"github.com/davide/ModRepository/routes/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {

	dsn := "host=192.168.0.113 user=davide password=SP589a%6 dbname=mod-repo-2 port=5432 sslmode=disable TimeZone=Europe/Rome"
	dbase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	carRepo := postgresrepo.CarRepositoryImpl{Db: dbase}
	trackRepo := postgresrepo.TrackRepositoryImpl{Db: dbase}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.113:27017").SetAuth(options.Credential{
		Username: "mongoAdmin",
		Password: "SP589a%6",
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

	tracksCollection := client.Database("mod_repo").Collection("tracks")
	carsCollection := client.Database("mod_repo").Collection("cars")

	web := routes.Web{
		CarHandler:    handlers.CarsHandlerImpl{CarCtrl: controllers.CarControllerImpl{Repo: carRepo}},
		TracksHandler: handlers.TrackHandlerImpl{TrackCtrl: controllers.TrackControllerImpl{Repo: trackRepo}},
		NationHandler: handlers.NationsHandlerImpl{CtrlNations: controllers.NationControllerImpl{Repo: mongorepo.NationsRepositoryImpl{CarsCollection: carsCollection, TracksCollection: tracksCollection}}},
		BrandsHandler: handlers.BrandsHandlerImpl{BrandCtrl: controllers.BrandControllerImpl{Repo: mongorepo.BrandRepositoryImpl{CarCollection: carsCollection}}},
	}
	web.Listen("6316")
}
