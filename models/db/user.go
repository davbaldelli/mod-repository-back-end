package db

type User struct {
	Username string
	Password string
	Role     string
	Salt     string
}
