package service

import (
	entity "sport_app"
	"sport_app/pkg/repository"
)

type StatsService struct {
	repo repository.Stats
}

func NewStatsService(repo repository.Stats) *StatsService {
	return &StatsService{repo: repo}
}

func (s *StatsService) GetActivityStateStats(visitorId int, activityName string) ([]entity.ActivityState, error){
	return s.repo.GetActivityStateStats(visitorId, activityName)
}
