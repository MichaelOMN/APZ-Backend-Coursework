package repository

import (
	"fmt"
	entity "sport_app"
	"strings"

	"github.com/jmoiron/sqlx"
)

type StatesTypesPostgres struct {
	db *sqlx.DB
}

func NewStatesTypesPostgres(db *sqlx.DB) *StatesTypesPostgres {
	return &StatesTypesPostgres{db}
}

func (a *StatesTypesPostgres) Create(ac entity.StatesTypes) (int, error) {
	query := `insert into StatesTypes(type_name, type_unit, type_desc) values($1, $2, $3) returning id`
	transaction, err := a.db.Begin()
	if err != nil {
		return -1, err
	}

	row := transaction.QueryRow(query, ac.Name, ac.Unit, ac.Decription)
	var id int
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, transaction.Commit()
}

func (a *StatesTypesPostgres) GetAll() ([]entity.StatesTypes, error) {
	return []entity.StatesTypes{}, nil
}

func (a *StatesTypesPostgres) GetById(acid int) (entity.StatesTypes, error) {
	query := `select id, type_name, type_unit, type_desc from StatesTypes where id = $1`
	var st entity.StatesTypes
	if err := a.db.Get(&st, query, acid); err != nil {
		return st, err
	}

	return st, nil
}

func (a *StatesTypesPostgres) Delete(acid int) error {
	return nil
}

func (a *StatesTypesPostgres) Update(stId int, ac entity.StatesTypesUpdateForm) error {
	stmts := make([]string, 0, 5)
	args := make([]any, 0, 5)
	argId := 1

	if ac.Name != nil {
		stmts = append(stmts, fmt.Sprintf("type_name=$%d", argId))
		args = append(args, *ac.Name)
		argId++
	}

	if ac.Decription != nil {
		stmts = append(stmts, fmt.Sprintf("type_desc=$%d", argId))
		args = append(args, *ac.Decription)
		argId++
	}

	if ac.Unit != nil {
		stmts = append(stmts, fmt.Sprintf("type_unit=$%d", argId))
		args = append(args, *ac.Unit)
		argId++
	}

	stmtsQuery := strings.Join(stmts, ", ")
	query := fmt.Sprintf(`update StatesTypes set %s where id = $%d`, stmtsQuery, argId)
	args = append(args, stId)

	_, err := a.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
