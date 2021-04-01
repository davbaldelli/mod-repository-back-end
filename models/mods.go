package models

type Mod struct {
	DownloadLink string
}

type Car struct {
	Mod
	Brand CarBrand
	ModelName string
	Categories []string
}


type CarBrand struct {
	Name string
	Nation Nation
}

type Nation struct {
	Name string
}

type Track struct {
	Mod
	layouts []Layout
	location Location
}

type Location struct {
	locationName string
	nation Nation
}

type Layout struct {
	lengthKm float32
	layoutType string
}