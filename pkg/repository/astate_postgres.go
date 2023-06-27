package repository

import (
	entity "sport_app"

	"github.com/jmoiron/sqlx"
)

type ActivityStatePostgres struct {
	db *sqlx.DB
}

func NewActivityStatePostgres(db *sqlx.DB) *ActivityStatePostgres {
	return &ActivityStatePostgres{db}
}

func (asp *ActivityStatePostgres) Create(as entity.ActivityState) (int, error) {
	tx, err := asp.db.Begin()
	if err != nil {
		return -1, err
	}

	query := `insert into ActivityStates(state_type_id, unit_amount, activity_name, duration_secs)
	values($1, $2, $3, $4) returning id`

	row := tx.QueryRow(query, as.StateTypeId, as.UnitAmount, as.ActivityName, as.Secs)
	//logrus.Fatalf("%s", as.ActivityName)
	var id int
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, tx.Commit()
}

func (asp *ActivityStatePostgres) GetAll() ([]entity.ActivityState, error) {
	return []entity.ActivityState{}, nil
}

func (asp *ActivityStatePostgres) GetById(activityId int) (entity.ActivityState, error) {
	query := `select id, state_type_id, unit_amount, at_datetime, activity_name, duration_secs
	from ActivityStates where id = $1`

	var as entity.ActivityState
	err := asp.db.Get(&as, query, activityId)
	if err != nil {
		return as, err
	}

	return as, nil
}

func (asp *ActivityStatePostgres) Delete(activityId int) error {
	return nil
}
