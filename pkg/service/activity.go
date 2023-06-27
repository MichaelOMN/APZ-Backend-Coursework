package service

import (
	entity "sport_app"
	"sport_app/pkg/repository"
)

type ActivityService struct {
	repo repository.Activity
}

func NewActivityService(repos repository.Activity) *ActivityService {
	return &ActivityService{repos}
}

func (ac *ActivityService) Create(a entity.Activity) (int, error) {
	return ac.repo.Create(a)
}

func (ac *ActivityService) GetAll() ([]entity.Activity, error) {
	return ac.repo.GetAll()
}

func (ac *ActivityService) GetById(aid int) (entity.Activity, error) {
	return ac.repo.GetById(aid)
}

func (ac *ActivityService) Update(acId int, form entity.ActivityUpdateForm) error {
	return ac.repo.Update(acId, form)
}

func (ac *ActivityService) Delete(acId int) error {
	return ac.repo.Delete(acId)
}
