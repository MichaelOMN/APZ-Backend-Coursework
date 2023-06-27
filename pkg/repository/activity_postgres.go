package repository

import (
	entity "sport_app"

	"github.com/jmoiron/sqlx"
)

type ActivityPostgres struct {
	db *sqlx.DB
}

func NewActivityPostgres(db *sqlx.DB) *ActivityPostgres {
	return &ActivityPostgres{db}
}

func (a *ActivityPostgres) Create(ac entity.Activity) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return -1, err
	}

	query := `insert into Activities(activity_name, activity_desc, club_id)
	values($1, $2, $3) returning id`

	var id int
	row := tx.QueryRow(query, ac.Name, ac.Description, ac.ClubId)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, tx.Commit()
}

func (a *ActivityPostgres) GetAll() ([]entity.Activity, error) {
	return []entity.Activity{}, nil
}

func (a *ActivityPostgres) GetById(acid int) (entity.Activity, error) {
	query := `select id, activity_name, activity_desc, club_id
	from Activities where id = $1`

	var act entity.Activity
	if err := a.db.Get(&act, query, acid); err != nil {
		return act, err
	}

	return act, nil
}

func (a *ActivityPostgres) Delete(acid int) error {
	query := `delete from Activities where id = $1`

	if _, err := a.db.Exec(query, acid); err != nil {
		return err
	}
	return nil
}

func (a *ActivityPostgres) Update(acId int, ac entity.ActivityUpdateForm) error {
	return nil
}
