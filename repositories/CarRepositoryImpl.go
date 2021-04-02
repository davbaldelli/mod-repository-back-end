package repositories

import (
	"github.com/davide/ModRepository/models"
)

type CarRepositoryImpl struct {
}

func (c CarRepositoryImpl) AddNewCar(car models.Car) error {
	return nil
}

func (c CarRepositoryImpl) GetAllCars() []models.Car {
	return []models.Car{
		{
			Mod:        models.Mod{DownloadLink: "FuckYou.org"},
			Brand:      models.CarBrand{Nation: models.Nation{Name: "Japan"}, Name: "Mazda"},
			ModelName:  "RX7 FD3S",
			Categories: []string{"JDM"},
		},
		{
			Mod:        models.Mod{DownloadLink: "FuckYou.org"},
			Brand:      models.CarBrand{Nation: models.Nation{Name: "Italy"}, Name: "Ferrari"},
			ModelName:  "F40",
			Categories: []string{"Italian Classic"},
		},
	}
}

func (c CarRepositoryImpl) GetCarsByNation(nation string) []models.Car {
	if nation == "Japan" {
		return []models.Car{
			{
				Mod:        models.Mod{DownloadLink: "FuckYou.org"},
				Brand:      models.CarBrand{Nation: models.Nation{Name: "Japan"}, Name: "Mazda"},
				ModelName:  "RX7 FD3S",
				Categories: []string{"JDM"},
			},
		}
	}
	if nation == "Italy" {
		return []models.Car{
			{
				Mod:        models.Mod{DownloadLink: "FuckYou.org"},
				Brand:      models.CarBrand{Nation: models.Nation{Name: "Italy"}, Name: "Ferrari"},
				ModelName:  "F40",
				Categories: []string{"Italian Classic"},
			},
		}
	}
	return []models.Car{}
}

func (c CarRepositoryImpl) GetCarByModel(model string) []models.Car {
	if model == "RX7 FD3S" {
		return []models.Car{{
			Mod:        models.Mod{DownloadLink: "FuckYou.org"},
			Brand:      models.CarBrand{Nation: models.Nation{Name: "Japan"}, Name: "Mazda"},
			ModelName:  "RX7 FD3S",
			Categories: []string{"JDM"},
		}}
	}
	if model == "F40" {
		return []models.Car{{

			Mod:        models.Mod{DownloadLink: "FuckYou.org"},
			Brand:      models.CarBrand{Nation: models.Nation{Name: "Italy"}, Name: "Ferrari"},
			ModelName:  "F40",
			Categories: []string{"Italian Classic"},
		}}
	}
	return []models.Car{}
}

func (c CarRepositoryImpl) GetCarsByBrand(s string) []models.Car {
	panic("implement me")
}

func (c CarRepositoryImpl) GetCarsByType(s string) []models.Car {
	panic("implement me")
}
