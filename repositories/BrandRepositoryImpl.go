package repositories

import (
	"context"
	"errors"
	"github.com/davide/ModRepository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BrandRepositoryImpl struct {
	CarCollection *mongo.Collection
}

func (b BrandRepositoryImpl) AddNewBrand(brand models.CarBrand) error {
	/*ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := b.CarCollection.InsertOne(ctx, brand)
	if err != nil {
		return err
	}
	return nil*/
	return errors.New("method not implemented")
}

func (b BrandRepositoryImpl) GetAllBrands() []models.CarBrand {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var brands []models.CarBrand

	values, err := b.CarCollection.Distinct(ctx, "brand", bson.D{{}})
	if err != nil {
		panic(err)
	}

	for _, value := range values {
		brands = append(brands, models.CarBrand{
			Name:   value.(bson.D).Map()["name"].(string),
			Nation: models.Nation{Name: value.(bson.D).Map()["nation"].(bson.D).Map()["name"].(string)},
		})
	}

	return brands
}

func (b BrandRepositoryImpl) GetBrandByNation(nationName string) []models.CarBrand {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var brands []models.CarBrand

	values, err := b.CarCollection.Distinct(ctx, "brand", bson.M{"brand.nation.name": nationName})
	if err != nil {
		panic(err)
	}

	for _, value := range values {
		brands = append(brands, models.CarBrand{
			Name:   value.(bson.D).Map()["name"].(string),
			Nation: models.Nation{Name: value.(bson.D).Map()["nation"].(bson.D).Map()["name"].(string)},
		})
	}

	return brands
}

func (b BrandRepositoryImpl) GetBrandByName(name string) []models.CarBrand {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var brands []models.CarBrand

	values, err := b.CarCollection.Distinct(ctx, "brand", bson.M{"brand.name": name})
	if err != nil {
		panic(err)
	}

	for _, value := range values {
		brands = append(brands, models.CarBrand{
			Name:   value.(bson.D).Map()["name"].(string),
			Nation: models.Nation{Name: value.(bson.D).Map()["nation"].(bson.D).Map()["name"].(string)},
		})
	}

	return brands
}
