package models

type Mod struct {
	DownloadLink string
}

type Car struct {
	Mod
	Brand      CarBrand
	ModelName  string
	Categories []string
}

type CarBrand struct {
	Name   string
	Nation Nation
}

type Nation struct {
	Name string
}

type Track struct {
	Mod
	Name     string
	Layouts  []Layout
	Location Location
}

type Location struct {
	LocationName string
	Nation       Nation
}

type Layout struct {
	Name     string
	LengthKm float32
	Type     string
}
