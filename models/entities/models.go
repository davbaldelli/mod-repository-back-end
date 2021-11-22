package entities

import "time"

type Drivetrain string

type Role string

func ToTrackTags(strings []string) []TrackTag {
	var tags []TrackTag
	for _, s := range strings {
		tags = append(tags, TrackTag(s))
	}
	return tags
}

type Transmission string

type TrackTag string

type LayoutType string

const (
	Admin   Role = "admin"
	Premium Role = "premium"
	Base    Role = "base"
)

const (
	F1          TrackTag = "F1"
	Touge       TrackTag = "Touge"
	NASCAR      TrackTag = "NASCAR"
	Rally       TrackTag = "Rally"
	CityTrack   TrackTag = "City Track"
	StreetTrack TrackTag = "Street Track"
	Fictional   TrackTag = "Fictional"
	Endurance   TrackTag = "Endurance"
	Drift       TrackTag = "Drift"
	Historic    TrackTag = "Historic"
	OpenWorld   TrackTag = "Open World"
	Karting     TrackTag = "Karting"
)

const (
	RoadCourse LayoutType = "Road Course"
	Oval       LayoutType = "Oval"
	AToB       LayoutType = "A to B"
)

const (
	RearWheelDrive  Drivetrain = "RWD"
	FrontWheelDrive Drivetrain = "FWD"
	AllWheelDrive   Drivetrain = "AWD"
)

const (
	Sequential Transmission = "Sequential"
	Manual     Transmission = "Manual"
)

type Mod struct {
	Id uint
	DownloadLink string
	Premium      bool
	Image        string
	Author       Author
}

type Author struct {
	Name string
	Link string
}

type Car struct {
	Mod
	Brand        CarBrand
	ModelName    string
	Categories   []CarCategory
	Year         uint
	Drivetrain   Drivetrain
	Transmission Transmission
	BHP          uint
	TopSpeed     uint
	Weight       uint
	Torque       uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CarCategory struct {
	Name string
}

type CarBrand struct {
	Name   string
	Nation Nation
}

type Nation struct {
	Name string
	Code string
}

type Track struct {
	Mod
	Name     string
	Tags     []TrackTag
	Layouts  []Layout
	Location string
	Nation   Nation
	Year     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Layout struct {
	Name     string
	LengthM  float32
	Category LayoutType
}

type User struct {
	Username string
	Password string
	Role     Role
}

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Username    string `json:"username"`
	TokenString string `json:"token"`
}
