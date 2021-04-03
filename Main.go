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

	web := routes.Web{
		CarHandler:    handlers.CarsHandlerImpl{CarCtrl: controllers.CarControllerImpl{Repo: repositories.CarRepositoryImpl{CarCollection: client.Database("mod_repo").Collection("cars")}}},
		TracksHandler: handlers.TrackHandlerImpl{TrackCtrl: controllers.TrackControllerImpl{Repo: repositories.TrackRepositoryImpl{TrackCollection: client.Database("mod_repo").Collection("tracks")}}},
	}
	web.Listen("8080")
}
