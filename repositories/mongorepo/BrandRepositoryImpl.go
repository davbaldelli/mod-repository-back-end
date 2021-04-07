package mongorepo

import (
	"context"
	"errors"
	"github.com/davide/ModRepository/models/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BrandRepositoryImpl struct {
	CarCollection *mongo.Collection
}

func (b BrandRepositoryImpl) AddNewBrand(brand entities.CarBrand) error {
	/*ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := b.CarCollection.InsertOne(ctx, brand)
	if err != nil {
		return err
	}
	return nil*/
	return errors.New("method not implemented")
}

func (b BrandRepositoryImpl) GetAllBrands() []entities.CarBrand {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var brands []entities.CarBrand

	values, err := b.CarCollection.Distinct(ctx, "brand", bson.D{{}})
	if err != nil {
		panic(err)
	}

	for _, value := range values {
		brands = append(brands, entities.CarBrand{
			Name:   value.(bson.D).Map()["name"].(string),
			Nation: entities.Nation{Name: value.(bson.D).Map()["nation"].(bson.D).Map()["name"].(string)},
		})
	}

	return brands
}

func (b BrandRepositoryImpl) GetBrandByNation(nationName string) []entities.CarBrand {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var brands []entities.CarBrand

	values, err := b.CarCollection.Distinct(ctx, "brand", bson.M{"brand.nation.name": nationName})
	if err != nil {
		panic(err)
	}

	for _, value := range values {
		brands = append(brands, entities.CarBrand{
			Name:   value.(bson.D).Map()["name"].(string),
			Nation: entities.Nation{Name: value.(bson.D).Map()["nation"].(bson.D).Map()["name"].(string)},
		})
	}

	return brands
}

func (b BrandRepositoryImpl) GetBrandByName(name string) []entities.CarBrand {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var brands []entities.CarBrand

	values, err := b.CarCollection.Distinct(ctx, "brand", bson.M{"brand.name": name})
	if err != nil {
		panic(err)
	}

	for _, value := range values {
		brands = append(brands, entities.CarBrand{
			Name:   value.(bson.D).Map()["name"].(string),
			Nation: entities.Nation{Name: value.(bson.D).Map()["nation"].(bson.D).Map()["name"].(string)},
		})
	}

	return brands
}
