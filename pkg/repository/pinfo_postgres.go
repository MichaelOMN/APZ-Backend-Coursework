package repository

import (
	"fmt"
	entity "sport_app"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PhysicalInfoPostgres struct {
	db *sqlx.DB
}

func NewPhysicalInfoPostgres(db *sqlx.DB) *PhysicalInfoPostgres {
	return &PhysicalInfoPostgres{db}
}

func (a *PhysicalInfoPostgres) Create(ac entity.PhysicalInfo) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return -1, err
	}

	var id int
	q := `insert into PhysicalInfo (visitor_id, height, weight) values ($1, $2, $3) returning id`
	//logrus.f("visitor_id = %d, height = %f, weight = %f", ac.VisitorId, ac.Height, ac.Weight)
	row := tx.QueryRow(q, ac.VisitorId, ac.Height, ac.Weight)
	err = row.Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, tx.Commit()
}

func (a *PhysicalInfoPostgres) GetAllWithVisitorId(visitorId int) ([]entity.PhysicalInfo, error) {
	return []entity.PhysicalInfo{}, nil
}

func (a *PhysicalInfoPostgres) GetByIdAndVisitorId(phId, visitorId int) (entity.PhysicalInfo, error) {
	q := `select id, visitor_id, height, weight from PhysicalInfo where id = $1 and visitor_id = $2`
	var pi entity.PhysicalInfo
	err := a.db.Get(&pi, q, phId, visitorId)
	if err != nil {
		return pi, err
	}

	return pi, nil
}

func (a *PhysicalInfoPostgres) Delete(acid int, userId int) error {
	return nil
}

func (a *PhysicalInfoPostgres) Update(userId int, pinfoId int, ac entity.PhysicalInfoUpdateForm) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if ac.Height != nil {
		setValues = append(setValues, fmt.Sprintf("height=$%d", argId))
		args = append(args, *ac.Height)
		argId++
	}

	if ac.Weight != nil {
		setValues = append(setValues, fmt.Sprintf("weight=$%d", argId))
		args = append(args, *ac.Weight)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s where id=$%d and visitor_id=$%d",
		"PhysicalInfo", setQuery, argId, argId+1)

	args = append(args, pinfoId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := a.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
