package helpers

import (
	"github.com/davide/ModRepository/models"
)

func IsAdmin(role models.Role) bool {
	return role == models.Admin
}

func IsPremium(role models.Role) bool {
	return role == models.Premium || role == models.FSRTeam || role == models.Admin
}
