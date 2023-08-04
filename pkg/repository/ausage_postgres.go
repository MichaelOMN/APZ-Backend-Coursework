package repository

import (
	"database/sql"
	entity "sport_app"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ActivityUsagePostgres struct {
	db *sqlx.DB
}

func NewActivityUsagePostgres(db *sqlx.DB) *ActivityUsagePostgres {
	return &ActivityUsagePostgres{db}
}

func (a *ActivityUsagePostgres) Create(ac entity.ActivityUsage) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return -1, err
	}

	query := ``
	var row *sql.Row

	if ac.Start != "" && ac.End != "" {
		query = `insert into ActivityUsage(visitor_id, activity_name, usage_start_datetime, usage_end_datetime, training_id)
		values($1, $2, $3, $4, $5) returning id`
		row = tx.QueryRow(query, ac.VisitorId, ac.ActivityName, ac.Start, ac.End, ac.TrainingId)
	} else {
		query = `insert into ActivityUsage(visitor_id, activity_name, training_id)
		values($1, $2, $3) returning id`
		row = tx.QueryRow(query, ac.VisitorId, ac.ActivityName, ac.TrainingId)
	}

	var id int
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, tx.Commit()
}

func (a *ActivityUsagePostgres) GetAll(visitorId, activityId int) ([]entity.ActivityUsage, error) {
	return []entity.ActivityUsage{}, nil
}

func (a *ActivityUsagePostgres) GetById(visitorId int, activityName string) (entity.ActivityUsage, error) {
	query := `select id, visitor_id, activity_name, usage_start_datetime, usage_end_datetime, training_id
	from ActivityUsage where visitor_id = $1 and activity_name = $2 order by usage_end_datetime desc limit 1`
	var au entity.ActivityUsage
	logrus.Errorf("%s - name\n", activityName)
	err := a.db.Get(&au, query, visitorId, activityName)
	if err != nil {
		return au, err
	}
	return au, nil
}

func (a *ActivityUsagePostgres) GetByActUsageId(visitorId int, activityUsageId int) (entity.ActivityUsage, error) {
	query := `select id, visitor_id, activity_name, usage_start_datetime, usage_end_datetime, training_id from
	ActivityUsage where visitor_id = $1 and id = $2 order by usage_end_datetime desc limit 1`
	var au entity.ActivityUsage
	err := a.db.Get(&au, query, visitorId, activityUsageId)
	if err != nil {
		return au, err
	}
	return au, nil
}

func (a *ActivityUsagePostgres) Delete(visitorId, activityUsageId int) error {
	query := `delete from activityusage where visitor_id = $1 and id = $2`
	if _, err := a.db.Exec(query, visitorId, activityUsageId); err != nil {
		return err;
	}
	return nil
}

func (a *ActivityUsagePostgres) Update(visitorId int, ac entity.ActivityUsageUpdateForm) error {
	if err := ac.Validate(); err != nil {
		return err
	}
	return nil
}
