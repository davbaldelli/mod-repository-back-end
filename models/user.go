package models

type Role string

const (
	Admin   Role = "admin"
	Premium Role = "premium"
	Base    Role = "base"
	FSRTeam Role = "fsrteam"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
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
