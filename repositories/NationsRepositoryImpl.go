package repositories

import (
	"context"
	"github.com/davide/ModRepository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type NationsRepositoryImpl struct {
	TracksCollection *mongo.Collection
	CarsCollection   *mongo.Collection
}

func (n NationsRepositoryImpl) GetAllBrandsNations() []models.Nation {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var nations []models.Nation
	values, err := n.CarsCollection.Distinct(ctx, "brand.nation.name", bson.M{})
	if err != nil {
		panic(err)
	}
	for _, value := range values {
		nations = append(nations, models.Nation{Name: value.(string)})
	}
	return nations
}

func (n NationsRepositoryImpl) GetAllTrackNations() []models.Nation {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var nations []models.Nation
	values, err := n.TracksCollection.Distinct(ctx, "location.nation.name", bson.M{})
	if err != nil {
		panic(err)
	}
	for _, value := range values {
		nations = append(nations, models.Nation{Name: value.(string)})
	}
	return nations
}
