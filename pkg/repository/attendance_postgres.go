package repository

import (
	entity "sport_app"

	"github.com/jmoiron/sqlx"
)

type AttendancePostgres struct {
	db *sqlx.DB
}

func NewAttendancePostgres(db *sqlx.DB) *AttendancePostgres {
	return &AttendancePostgres{db}
}

func (a *AttendancePostgres) Create(ac entity.Attendance) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return -1, err
	}

	query := `insert into Attendance(visitor_id, training_id) values($1, $2) returning id`
	var id int
	row := tx.QueryRow(query, ac.VisitorId, ac.TrainingId)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, tx.Commit()
}

func (a *AttendancePostgres) GetAll(visitorId int) ([]entity.Attendance, error) {
	return []entity.Attendance{}, nil
}

func (a *AttendancePostgres) GetById(visitorId, acid int) (entity.Attendance, error) {
	query := `select id, visitor_id, training_id from Attendance where id = $1 and visitor_id = $2`
	var at entity.Attendance
	if err := a.db.Get(&at, query, acid, visitorId); err != nil {
		return at, err
	}
	return at, nil
}

func (a *AttendancePostgres) Delete(visitorId, acid int) error {
	return nil
}

func (a *AttendancePostgres) Update(visitorId int, ac entity.AttendanceUpdateForm) error {
	return nil
}
