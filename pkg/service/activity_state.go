package service

import (
	entity "sport_app"
	"sport_app/pkg/repository"
)

type ActivityStateService struct {
	repo repository.ActivityState
}

func NewActivityStateService(repos repository.ActivityState) *ActivityStateService {
	return &ActivityStateService{repos}
}

func (ac *ActivityStateService) Create(a entity.ActivityState) (int, error) {
	return ac.repo.Create(a)
}

func (ac *ActivityStateService) GetAll() ([]entity.ActivityState, error) {
	return ac.repo.GetAll()
}

func (ac *ActivityStateService) GetById(activityId int) (entity.ActivityState, error) {
	return ac.repo.GetById(activityId)
}

func (ac *ActivityStateService) Delete(activityId int) error {
	return ac.repo.Delete(activityId)
}
