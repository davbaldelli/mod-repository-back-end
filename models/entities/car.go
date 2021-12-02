package entities

type Drivetrain string
type CarType string
type Transmission string

const (
	EnduranceCar CarType = "Endurance"
	Formula              = "Formula"
	GT                   = "GT"
	Prototype            = "Prototype"
	RallyCar             = "Rally"
	Street               = "Street"
	Tuned                = "Tuned"
	Touring              = "Touring"
	Vintage              = "Vintage"
	StockCar             = "Stock Car"
)

const (
	RearWheelDrive  Drivetrain = "RWD"
	FrontWheelDrive            = "FWD"
	AllWheelDrive              = "AWD"
)

const (
	Sequential Transmission = "Sequential"
	Manual                  = "Manual"
)

type Car struct {
	Mod
	Brand        CarBrand      `json:"brand"`
	ModelName    string        `json:"modelName"`
	Categories   []CarCategory `json:"categories"`
	Year         uint          `json:"year"`
	Drivetrain   Drivetrain    `json:"drivetrain"`
	Transmission Transmission  `json:"transmission"`
	BHP          uint          `json:"bhp"`
	TopSpeed     uint          `json:"topSpeed"`
	Weight       uint          `json:"weight"`
	Torque       uint          `json:"torque"`
}

type CarCategory struct {
	Name CarType `json:"name"`
}
