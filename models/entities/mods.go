package entities

type Mod struct {
	DownloadLink string
}

type Car struct {
	Mod
	Brand      CarBrand
	ModelName  string
	Categories []CarCategory
}

type CarCategory struct {
	Name string
}

type TrackCategory struct {
	Name string
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
	Location string
	Nation   Nation
}

type Layout struct {
	Name     string
	LengthKm float32
	Category TrackCategory
}
