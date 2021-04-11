package presentation

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarPresentation struct {
	Model        string
	Brand        string
	DownloadLink string
}

type TrackPresentation struct {
	Name         string
	Location     string
	Nation       string
	DownloadLink string
}
type BrandPresentation struct {
	Name   string
	Nation string
}

type NationPresentation struct {
	Name string
}

type NationBrands struct {
	Nation string
	Brands []string
}

func ofNation(nation entities.Nation) NationPresentation {
	return NationPresentation{
		Name: nation.Name,
	}
}

func OfAllNations(nations []entities.Nation) []NationPresentation {
	var nationsPres []NationPresentation
	for _, nation := range nations {
		nationsPres = append(nationsPres, ofNation(nation))
	}
	return nationsPres
}

func ofBrand(brand entities.CarBrand) BrandPresentation {
	return BrandPresentation{
		Name:   brand.Name,
		Nation: brand.Nation.Name,
	}
}

func OfAllBrands(brands []entities.CarBrand) []BrandPresentation{
	var brandsPres []BrandPresentation
	for _, brand := range brands {
		brandsPres = append(brandsPres, ofBrand(brand))
	}
	return brandsPres
}

func OfAllBrandsGroupedByNation(brands []entities.CarBrand) []NationBrands {
	brandsMap := make(map[string][]string)
	for _, brand := range brands {
		brandsMap[brand.Nation.Name] = append(brandsMap[brand.Nation.Name], brand.Name)
	}
	var nationsBrands []NationBrands
	for nation, nBrands := range brandsMap{
		nationsBrands = append(nationsBrands, NationBrands{Nation: nation, Brands: nBrands})
	}
	return nationsBrands
}

func ofCars(car entities.Car) CarPresentation {
	return CarPresentation{
		DownloadLink: car.DownloadLink,
		Brand:        car.Brand.Name,
		Model:        car.ModelName,
	}
}

func OfAllCars(cars []entities.Car) []CarPresentation {
	var presCars []CarPresentation
	for _, car := range cars {
		presCars = append(presCars, ofCars(car))
	}
	return presCars
}

func ofTrack(track entities.Track) TrackPresentation {
	return TrackPresentation{
		Name:         track.Name,
		DownloadLink: track.DownloadLink,
		Location:     track.Location,
		Nation:       track.Nation.Name,
	}
}
func OfAllTracks(tracks []entities.Track) []TrackPresentation {
	var tracksPres []TrackPresentation
	for _, track := range tracks {
		tracksPres = append(tracksPres, ofTrack(track))
	}
	return tracksPres
}
