package repositories

import (
	"context"
	"fmt"
	"github.com/davide/ModRepository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CarRepositoryImpl struct {
	CarCollection *mongo.Collection
}

func (c CarRepositoryImpl) AddNewCar(car models.Car) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := c.CarCollection.InsertOne(ctx, car)
	if err != nil {
		return err
	}
	return nil
}

func (c CarRepositoryImpl) GetAllCars() []models.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []models.Car
	cursor, err := c.CarCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &cars); err != nil {
		panic(err)
	}

	values, _ := c.CarCollection.Distinct(ctx, "brand.nation.name", bson.M{})

	fmt.Println(values)

	return cars
}

func (c CarRepositoryImpl) GetCarsByNation(nation string) []models.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []models.Car
	cursor, err := c.CarCollection.Find(ctx, bson.M{"brand.nation.name": nation})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &cars); err != nil {
		panic(err)
	}
	return cars
}

func (c CarRepositoryImpl) GetCarByModel(model string) []models.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []models.Car
	cursor, err := c.CarCollection.Find(ctx, bson.M{"modelname": model})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &cars); err != nil {
		panic(err)
	}
	return cars
}

func (c CarRepositoryImpl) GetCarsByBrand(brandName string) []models.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []models.Car
	cursor, err := c.CarCollection.Find(ctx, bson.M{"brand.name": brandName})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &cars); err != nil {
		panic(err)
	}
	return cars
}

func (c CarRepositoryImpl) GetCarsByType(category string) []models.Car {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []models.Car
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
