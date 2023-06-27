package repository

import (
	entity "sport_app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateVisitor(user entity.Visitor) (int, error)
	GetVisitor(username, password string) (entity.Visitor, error)
	CreateCoach(user entity.Coach) (int, error)
	GetCoach(username, password string) (entity.Coach, error)
}

type Club interface {
	Create(club entity.Club) (int, error)
	GetAll() ([]entity.Club, error)
	GetById(clubId int) (entity.Club, error)
	Delete(clubId int) error
	Update(clubId int, club entity.ClubUpdateForm) error
}

type PhysicalInfo interface {
	Create(ph entity.PhysicalInfo) (int, error)
	GetAllWithVisitorId(visitorId int) ([]entity.PhysicalInfo, error)
	GetByIdAndVisitorId(phId, visitorId int) (entity.PhysicalInfo, error)
	Delete(phId int, userId int) error
	Update(visitorId int, pinfoId int, ph entity.PhysicalInfoUpdateForm) error
}

type Training interface {
	Create(t entity.Training) (int, error)
	GetAll() ([]entity.Training, error)
	GetById(tid int) (entity.Training, error)
	Delete(tid int) error
	// Update(t entity.Training) error
}

type StatesTypes interface {
	Create(st entity.StatesTypes) (int, error)
	GetAll() ([]entity.StatesTypes, error)
	GetById(stid int) (entity.StatesTypes, error)
	Delete(stid int) error
	Update(stId int, st entity.StatesTypesUpdateForm) error
}

type Activity interface {
	Create(a entity.Activity) (int, error)
	GetAll() ([]entity.Activity, error)
	GetById(aid int) (entity.Activity, error)
	Delete(aid int) error
	Update(acId int, a entity.ActivityUpdateForm) error
}

type Attendance interface {
	Create(a entity.Attendance) (int, error)
	GetAll(visitorId int) ([]entity.Attendance, error)
	GetById(visitorId, aid int) (entity.Attendance, error)
	Delete(visitorId, aid int) error
	Update(visitorId int, a entity.AttendanceUpdateForm) error
}

type ActivityUsage interface {
	Create(au entity.ActivityUsage) (int, error)
	GetAll(visitorId, activityId int) ([]entity.ActivityUsage, error)
	GetById(visitorId int, activityName string) (entity.ActivityUsage, error)
	GetByActUsageId(visitorId int, activityUsageId int) (entity.ActivityUsage, error)
	Delete(visitorId, activityId int) error
	Update(visitorId int, au entity.ActivityUsageUpdateForm) error
}

type PhysicalState interface {
	Create(ps entity.PhysicalState) (int, error)
	GetAllByVisitorId(visitorId int) ([]entity.PhysicalState, error)
	GetByVisitorId(visitorId, phstId int) (entity.PhysicalState, error)
	Delete(visitorId, psid int) error
	//Update(visitorId, ps entity.PhysicalStateUpdateForm) error
}

type ActivityState interface {
	Create(as entity.ActivityState) (int, error)
	GetAll() ([]entity.ActivityState, error)
	GetById(activityId int) (entity.ActivityState, error)
	Delete(activityId int) error
	//Update(as entity.ActivityState) error
}

type Repository struct {
	Authorization
	Club
	PhysicalInfo
	Training
	StatesTypes
	Activity
	Attendance
	ActivityUsage
	PhysicalState
	ActivityState
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Club:          NewClubPostgres(db),
		PhysicalInfo:  NewPhysicalInfoPostgres(db),
		Training:      NewTrainingPostgres(db),
		StatesTypes:   NewStatesTypesPostgres(db),
		Activity:      NewActivityPostgres(db),
		Attendance:    NewAttendancePostgres(db),
		ActivityUsage: NewActivityUsagePostgres(db),
		PhysicalState: NewPhysicalStatePostgres(db),
		ActivityState: NewActivityStatePostgres(db),
	}
}
