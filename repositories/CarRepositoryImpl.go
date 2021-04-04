package repositories

import (
	"context"
	"github.com/davide/ModRepository/models/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CarRepositoryImpl struct {
	CarCollection *mongo.Collection
}

func (c CarRepositoryImpl) AddNewCar(car entities.Car) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := c.CarCollection.InsertOne(ctx, car)
	if err != nil {
		return err
	}
	return nil
}

func (c CarRepositoryImpl) GetAllCars() []entities.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []entities.Car
	cursor, err := c.CarCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &cars); err != nil {
		panic(err)
	}

	return cars
}

func (c CarRepositoryImpl) GetCarsByNation(nation string) []entities.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []entities.Car
	cursor, err := c.CarCollection.Find(ctx, bson.M{"brand.nation.name": nation})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &cars); err != nil {
		panic(err)
	}
	return cars
}

func (c CarRepositoryImpl) GetCarByModel(model string) []entities.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []entities.Car
	cursor, err := c.CarCollection.Find(ctx, bson.M{"modelname": model})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &cars); err != nil {
		panic(err)
	}
	return cars
}

func (c CarRepositoryImpl) GetCarsByBrand(brandName string) []entities.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []entities.Car
	cursor, err := c.CarCollection.Find(ctx, bson.M{"brand.name": brandName})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &cars); err != nil {
		panic(err)
	}
	return cars
}

func (c CarRepositoryImpl) GetCarsByType(category string) []entities.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []entities.Car
	categoryArr := []string{category}
	cursor, err := c.CarCollection.Find(ctx, bson.M{"categories": bson.M{"$in": categoryArr}})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &cars); err != nil {
		panic(err)
	}
	return cars
}
