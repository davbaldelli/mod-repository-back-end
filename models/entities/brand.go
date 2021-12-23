package entities

type CarBrand struct {
	Name   string `json:"name"`
	Nation Nation `json:"nation"`
	Logo string `json:"logo"`
}

