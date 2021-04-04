package main

import (
	"context"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/repositories"
	"github.com/davide/ModRepository/routes"
	"github.com/davide/ModRepository/routes/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {

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
		CarHandler:    handlers.CarsHandlerImpl{CarCtrl: controllers.CarControllerImpl{Repo: repositories.CarRepositoryImpl{CarCollection: carsCollection}}},
		TracksHandler: handlers.TrackHandlerImpl{TrackCtrl: controllers.TrackControllerImpl{Repo: repositories.TrackRepositoryImpl{TrackCollection: tracksCollection}}},
		NationHandler: handlers.NationsHandlerImpl{CtrlNations: controllers.NationControllerImpl{Repo: repositories.NationsRepositoryImpl{CarsCollection: carsCollection, TracksCollection: tracksCollection}}},
		BrandsHandler: handlers.BrandsHandlerImpl{BrandCtrl: controllers.BrandControllerImpl{Repo: repositories.BrandRepositoryImpl{CarCollection: carsCollection}}},
	}
	web.Listen("6316")
}
