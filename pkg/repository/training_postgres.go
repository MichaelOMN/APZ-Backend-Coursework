package repository

import (
	entity "sport_app"

	"github.com/jmoiron/sqlx"
)

type TrainingPostgres struct {
	db *sqlx.DB
}

func NewTrainingPostgres(db *sqlx.DB) *TrainingPostgres {
	return &TrainingPostgres{db}
}

func (a *TrainingPostgres) Create(ac entity.Training) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return -1, err
	}
	var id int
	query := `insert into Trainings(start_datetime, end_datetime, coach_id, club_id) values($1, $2, $3, $4) returning id`
	row := tx.QueryRow(query, ac.Start, ac.End, ac.CoachId, ac.ClubId)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, tx.Commit()
}

func (a *TrainingPostgres) GetAll() ([]entity.Training, error) {
	query := `select * from trainings`
	var ts []entity.Training

	err := a.db.Select(&ts, query)

	return ts, err
}

func (a *TrainingPostgres) GetById(acid int) (entity.Training, error) {
	query := `select id, start_datetime, end_datetime, coach_id, club_id from Trainings
	where id = $1`
	var t entity.Training

	if err := a.db.Get(&t, query, acid); err != nil {
		return t, err
	}
	return t, nil
}

func (a *TrainingPostgres) Delete(acid int) error {
	return nil
}

func (a *TrainingPostgres) Update(ac entity.Training) error {
	return nil
}
