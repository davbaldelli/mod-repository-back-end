package entities

type Drivetrain string

type GearType string

const (
	RearWheelDrive Drivetrain = "RWD"
	FrontWheelDrive = "FWD"
	AllWheelDrive = "AWD"
)

const (
	Sequential GearType = "SEQUENTIAL"
	Manual = "MANUAL"
)

type Mod struct {
	DownloadLink string
}

type Car struct {
	Mod
	Brand      CarBrand
	ModelName  string
	Categories []CarCategory
	Year uint
	Drivetrain Drivetrain
	GearType GearType
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
	Year uint
}

type Layout struct {
	Name     string
	LengthM  float32
	Category TrackCategory
}

type User struct {
	Username string
	Password string
	IsAdmin bool
}