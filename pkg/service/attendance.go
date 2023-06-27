package service

import (
	entity "sport_app"
	"sport_app/pkg/repository"
)

type AttendanceService struct {
	repo repository.Attendance
}

func NewAttendanceService(repo repository.Attendance) *AttendanceService {
	return &AttendanceService{repo}
}

func (at *AttendanceService) Create(a entity.Attendance) (int, error) {
	return at.repo.Create(a)
}

func (at *AttendanceService) GetAll(visitorId int) ([]entity.Attendance, error) {
	return at.repo.GetAll(visitorId)
}

func (at *AttendanceService) GetById(visitorId, activityId int) (entity.Attendance, error) {
	return at.repo.GetById(visitorId, activityId)
}

func (at *AttendanceService) Update(visitorId int, form entity.AttendanceUpdateForm) error {
	if err := form.Validate(); err != nil {
		return err
	}
	return at.repo.Update(visitorId, form)
}

func (at *AttendanceService) Delete(visitorId int, atId int) error {
	return at.repo.Delete(visitorId, atId)
}
