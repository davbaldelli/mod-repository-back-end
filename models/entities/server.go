package entities

type Server struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	JoinLink    string `json:"joinLink"`
	Password    string `json:"password"`
	Online      bool   `json:"online"`
	Track       uint   `json:"track"`
	Cars        []uint `json:"cars"`
	OutsideCars []OutsideMod `json:"outsideCars"`
}

type OutsideMod struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DownloadLink string `json:"downloadLink"`
	ServerId uint `json:"serverId"`
}