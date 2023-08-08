package repository

import (
	entity "sport_app"

	"github.com/jmoiron/sqlx"
)

type StatsPostgres struct {
	db *sqlx.DB
}

func NewStatsPostgres(db *sqlx.DB) *StatsPostgres {
	return &StatsPostgres{db}
}

func (a *StatsPostgres) GetActivityStateStats(visitorId int, activityName string) ([]entity.ActivityState, error) {
	query := `select ast.* from activitystates ast
	join activityusage aus on ast.activity_name = aus.activity_name AND
		ast.at_datetime between aus.usage_start_datetime and aus.usage_end_datetime
	where aus.visitor_id = $1 AND aus.activity_name = $2;`

	var raw []entity.ActivityState;
	if err := a.db.Select(&raw, query, visitorId, activityName); err != nil {
		return raw, err
	}
	return raw, nil
}
