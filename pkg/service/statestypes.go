package service

import (
	entity "sport_app"
	rep "sport_app/pkg/repository"
)

type StatesTypesService struct {
	repo rep.StatesTypes
}

func NewStatesTypesService(repo rep.StatesTypes) *StatesTypesService {
	return &StatesTypesService{repo}
}

func (a *StatesTypesService) Create(ac entity.StatesTypes) (int, error) {
	return a.repo.Create(ac)
}

func (a *StatesTypesService) GetAll() ([]entity.StatesTypes, error) {
	return a.repo.GetAll()
}

func (a *StatesTypesService) GetById(acid int) (entity.StatesTypes, error) {
	return a.repo.GetById(acid)
}

func (a *StatesTypesService) Delete(acid int) error {
	return a.repo.Delete(acid)
}

func (a *StatesTypesService) Update(stId int, ac entity.StatesTypesUpdateForm) error {
	return a.repo.Update(stId, ac)
}
