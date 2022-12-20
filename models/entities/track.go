package entities

type TrackTag string

type LayoutType string

const (
	F1          TrackTag = "F1"
	Touge       TrackTag = "Touge"
	NASCAR      TrackTag = "NASCAR"
	RallyTrack  TrackTag = "Rally"
	StreetTrack TrackTag = "Street Track"
	Fictional   TrackTag = "Fictional"
	Drift       TrackTag = "Drift"
	Historic    TrackTag = "Historic"
	OpenWorld   TrackTag = "Open World"
	Karting     TrackTag = "Karting"
	LaserScan   TrackTag = "Laser Scan"
)

const (
	RoadCourse LayoutType = "Road Course"
	Oval       LayoutType = "Oval"
	AToB       LayoutType = "A to B"
)

type Track struct {
	Mod
	Name     string     `json:"name"`
	Tags     []TrackTag `json:"tags"`
	Layouts  []Layout   `json:"layouts"`
	Location string     `json:"location"`
	Nation   Nation     `json:"nation"`
	Year     uint       `json:"year"`
}

type Layout struct {
	Name     string     `json:"name"`
	LengthM  float32    `json:"lengthM"`
	Category LayoutType `json:"category"`
}
