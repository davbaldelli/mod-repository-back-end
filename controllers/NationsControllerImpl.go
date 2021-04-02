package controllers

import (
	"errors"
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type NationControllerImpl struct {
	repo repositories.NationRepository
}

type nationParamCondition func(models.Nation) bool

func (n NationControllerImpl) getAllNations() []models.Nation {
	return n.repo.GetAllNations()
}

func (n NationControllerImpl) getNationByName(name string) (models.Nation, error) {
	for _, nation := range n.repo.GetAllNations() {
		if nation.Name == name {
			return nation, nil
		}
	}
	return models.Nation{}, errors.New("nation" + name + " not found")
}

func (n NationControllerImpl) addNewNation(name string) error {
	return n.addNewNation(name)
}

func (n NationControllerImpl) nationsResearchByParam(ncp nationParamCondition) []models.Nation {
	var nations []models.Nation
	for _, nation := range n.getAllNations() {
		if ncp(nation) {
			nations = append(nations, nation)
		}
	}
	return nations
}
