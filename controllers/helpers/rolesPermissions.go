package helpers

import "github.com/davide/ModRepository/models/entities"

func IsAdmin(role entities.Role) bool {
	return role == entities.Admin
}

func IsPremium(role entities.Role) bool {
	return role == entities.Premium || role == entities.FSRTeam || role == entities.Admin
}
