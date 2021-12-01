package entities

import "time"

type Drivetrain string

type Role string

type Transmission string

type TrackTag string

type LayoutType string

type CarType string

const (
	Admin   Role = "admin"
	Premium Role = "premium"
	Base    Role = "base"
)

const (
	F1          TrackTag = "F1"
	Touge       TrackTag = "Touge"
	NASCAR      TrackTag = "NASCAR"
	RallyTrack       TrackTag = "Rally"
	CityTrack   TrackTag = "City Track"
	StreetTrack TrackTag = "Street Track"
	Fictional   TrackTag = "Fictional"
	EnduranceTrack   TrackTag = "Endurance"
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

const (
	EnduranceCar CarType = "Endurance"
	Formula = "Formula"
	GT = "GT"
	Prototype = "Prototype"
	RallyCar = "Rally"
	Street = "Street"
	Tuned = "Tuned"
	Touring = "Touring"
	Vintage = "Vintage"
	StockCar = "Stock Car"
)

type Mod struct {
	Id uint `json:"id"`
	DownloadLink string `json:"downloadLink"`
	Premium      bool `json:"premium"`
	Image        string `json:"image"`
	Author       Author `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Rating uint `json:"rating"`
	Version string `json:"version"`
}

type Author struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type Car struct {
	Mod
	Brand        CarBrand `json:"brand"`
	ModelName    string `json:"modelName"`
	Categories   []CarCategory `json:"categories"`
	Year         uint `json:"year"`
	Drivetrain   Drivetrain `json:"drivetrain"`
	Transmission Transmission `json:"transmission"`
	BHP          uint `json:"bhp"`
	TopSpeed     uint `json:"topSpeed"`
	Weight       uint `json:"weight"`
	Torque       uint `json:"torque"`
}

type CarCategory struct {
	Name CarType `json:"name"`
}

type CarBrand struct {
	Name   string `json:"name"`
	Nation Nation `json:"nation"`
}

type Nation struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Track struct {
	Mod
	Name     string `json:"name"`
	Tags     []TrackTag `json:"tags"`
	Layouts  []Layout `json:"layouts"`
	Location string `json:"location"`
	Nation   Nation `json:"nation"`
	Year     uint `json:"year"`
}

type Layout struct {
	Name     string `json:"name"`
	LengthM  float32 `json:"lengthM"`
	Category LayoutType `json:"category"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role `json:"role"`
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
