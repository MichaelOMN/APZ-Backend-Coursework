package service

import (
	entity "sport_app"
	"sport_app/pkg/repository"
)

type TrainingService struct {
	repo repository.Training
}

func NewTrainingService(repo repository.Training) *TrainingService {
	return &TrainingService{repo}
}

func (ts *TrainingService) Create(t entity.Training) (int, error) {
	return ts.repo.Create(t)
}

func (ts *TrainingService) GetById(tid int) (entity.Training, error) {
	return ts.repo.GetById(tid)
}

func (ts *TrainingService) GetAll() ([]entity.Training, error) {
	return ts.repo.GetAll()
}

func (ts *TrainingService) Delete(tid int) error {
	return ts.repo.Delete(tid)
}

// 	GetAll() ([]entity.Training, error)
// 	GetById(tid int) (entity.Training, error)
// 	Delete(tid int) error
// 	Update(t entity.Training) error
