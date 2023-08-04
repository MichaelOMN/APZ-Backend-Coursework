package service

import (
	entity "sport_app"
	"sport_app/pkg/repository"
)

type ActivityUsageService struct {
	repo repository.ActivityUsage
}

func NewActivityUsageService(repos repository.ActivityUsage) *ActivityUsageService {
	return &ActivityUsageService{repos}
}

func (ac *ActivityUsageService) Create(a entity.ActivityUsage) (int, error) {
	return ac.repo.Create(a)
}

func (ac *ActivityUsageService) GetAll(visitorId, activityId int) ([]entity.ActivityUsage, error) {
	return ac.repo.GetAll(visitorId, activityId)
}

func (ac *ActivityUsageService) GetById(visitorId int, activityName string) (entity.ActivityUsage, error) {
	return ac.repo.GetById(visitorId, activityName)
}

func (ac *ActivityUsageService) GetByActUsageId(visitorId int, activityUsageId int) (entity.ActivityUsage, error){
	return ac.repo.GetByActUsageId(visitorId, activityUsageId)
}

func (ac *ActivityUsageService) Update(visitorId int, form entity.ActivityUsageUpdateForm) error {
	return ac.repo.Update(visitorId, form)
}

func (ac *ActivityUsageService) Delete(visitorId, activityUsageId int) error {
	return ac.repo.Delete(visitorId, activityUsageId)
}
