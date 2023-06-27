package service

import (
	entity "sport_app"
	"sport_app/pkg/repository"
)

type PhysicalInfoService struct {
	repo repository.PhysicalInfo
}

func NewPhysicalInfoService(repo repository.PhysicalInfo) *PhysicalInfoService {
	return &PhysicalInfoService{repo}
}

func (p *PhysicalInfoService) Create(ep entity.PhysicalInfo) (int, error) {
	id, err := p.repo.Create(ep)
	return id, err
}

func (p *PhysicalInfoService) GetAllWithVisitorId(visitorId int) ([]entity.PhysicalInfo, error) {
	pis, err := p.repo.GetAllWithVisitorId(visitorId)
	return pis, err
}

func (p *PhysicalInfoService) GetByIdAndVisitorId(pinfoId, visitorId int) (entity.PhysicalInfo, error) {
	pi, err := p.repo.GetByIdAndVisitorId(pinfoId, visitorId)
	return pi, err
}

func (p *PhysicalInfoService) Delete(visitorId, pinfoId int) error {
	err := p.repo.Delete(pinfoId, visitorId)
	return err
}

func (p *PhysicalInfoService) Update(visitorId int, pinfoId int, form entity.PhysicalInfoUpdateForm) error {
	return p.repo.Update(visitorId, pinfoId, form)
}
