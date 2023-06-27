package repository

import (
	entity "sport_app"

	"github.com/jmoiron/sqlx"
)

type PhysicalStatePostgres struct {
	db *sqlx.DB
}

func NewPhysicalStatePostgres(db *sqlx.DB) *PhysicalStatePostgres {
	return &PhysicalStatePostgres{db}
}

func (a *PhysicalStatePostgres) Create(ac entity.PhysicalState) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return -1, err
	}

	query := `insert into PhysicalStates(state_type_id,unit_amount,at_datetime,activity_usage_id,duration_secs)
	values($1, $2, $3, $4, $5) returning id`

	row := tx.QueryRow(query, ac.StateTypeId, ac.UnitAmount, ac.At, ac.ActivityUsageId, ac.Secs)
	var id int
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, tx.Commit()
}

func (a *PhysicalStatePostgres) GetAllByVisitorId(visitorId int) ([]entity.PhysicalState, error) {
	return []entity.PhysicalState{}, nil
}

func (a *PhysicalStatePostgres) GetByVisitorId(visitorId int, phstId int) (entity.PhysicalState, error) {
	query := `select id,state_type_id,unit_amount,at_datetime,activity_usage_id,duration_secs
	from PhysicalStates where id = $1 and activity_usage_id in (select id from ActivityUsage
	where visitor_id = $2)`

	var ps entity.PhysicalState
	err := a.db.Get(&ps, query, phstId, visitorId)
	if err != nil {
		return ps, err
	}

	return ps, nil
}

func (a *PhysicalStatePostgres) Delete(visitorId, phstId int) error {
	return nil
}
