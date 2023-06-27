package service

import (
	entity "sport_app"
	"sport_app/pkg/repository"
)

type ClubService struct {
	repo repository.Club
}

func NewClubService(repo repository.Club) *ClubService {
	return &ClubService{repo}
}

func (c *ClubService) Create(club entity.Club) (int, error) {
	clubId, err := c.repo.Create(club)
	return clubId, err
}

func (c *ClubService) GetAll() ([]entity.Club, error) {
	clubs, err := c.repo.GetAll()
	return clubs, err
}

func (c *ClubService) GetById(clubId int) (entity.Club, error) {
	club, err := c.repo.GetById(clubId)
	return club, err
}

func (c *ClubService) Delete(clubId int) error {
	err := c.repo.Delete(clubId)
	return err
}

func (c *ClubService) Update(clubId int, updform entity.ClubUpdateForm) error {
	if err := updform.Validate(); err != nil {
		return err
	}
	return c.repo.Update(clubId, updform)
}
