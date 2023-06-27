package repository

import (
	"fmt"
	entity "sport_app"

	"github.com/jmoiron/sqlx"
)

type ClubPostgres struct {
	db *sqlx.DB
}

func NewClubPostgres(db *sqlx.DB) *ClubPostgres {
	return &ClubPostgres{db}
}

func (a *ClubPostgres) Create(ac entity.Club) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return -1, err
	}

	var id int
	query := fmt.Sprintf(`INSERT INTO %s (address, club_name) VALUES ($1, $2) RETURNING id`, clubsTable)
	row := tx.QueryRow(query, ac.Address, ac.Name)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return -1, err
	}

	return id, tx.Commit()
}

func (a *ClubPostgres) GetAll() ([]entity.Club, error) {
	var clubs []entity.Club
	query := fmt.Sprintf(`SELECT id, address, club_name FROM %s`, clubsTable)

	err := a.db.Select(&clubs, query)
	return clubs, err
}

func (a *ClubPostgres) GetById(acid int) (entity.Club, error) {
	var club entity.Club
	query := fmt.Sprintf(`SELECT id, address, club_name FROM %s WHERE id=$1`, clubsTable)

	err := a.db.Get(&club, query, acid)

	return club, err
}

func (a *ClubPostgres) Delete(acid int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, clubsTable)
	_, err := a.db.Exec(query, acid)

	return err
}

func (a *ClubPostgres) Update(clubId int, updform entity.ClubUpdateForm) error {
	return nil
}
