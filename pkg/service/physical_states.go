package service

import (
	entity "sport_app"
	"sport_app/pkg/repository"
)

type PhysicalStateService struct {
	repo repository.PhysicalState
}

func NewPhysicalStateService(repos repository.PhysicalState) *PhysicalStateService {
	return &PhysicalStateService{repos}
}

func (ac *PhysicalStateService) Create(a entity.PhysicalState) (int, error) {
	return ac.repo.Create(a)
}

func (ac *PhysicalStateService) GetAllByVisitorId(visitorId int) ([]entity.PhysicalState, error) {
	return ac.repo.GetAllByVisitorId(visitorId)
}

func (ac *PhysicalStateService) GetByVisitorId(visitorId, phstId int) (entity.PhysicalState, error) {
	return ac.repo.GetByVisitorId(visitorId, phstId)
}

func (ac *PhysicalStateService) Delete(visitorId, phstId int) error {
	return ac.repo.Delete(visitorId, phstId)
}
