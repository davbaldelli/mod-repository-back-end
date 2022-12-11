package entities

type Server struct {
	Id          uint
	Name        string
	Description string
	JoinLink    string
	Password    string
	Online      bool
	Track       uint
	Cars        []uint
}
