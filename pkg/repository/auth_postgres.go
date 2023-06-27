package repository

import (
	"errors"
	"fmt"
	entity "sport_app"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) CreateVisitor(visitor entity.Visitor) (int, error) {
	var has bool = false
	visitorExistsQuery := `select exists(select * from Visitors where name = $1) or
						exists(select * from Coaches where name = $1)`

	row := r.db.QueryRow(visitorExistsQuery, visitor.Name)
	if err := row.Scan(&has); err != nil {
		return -1, err
	}

	if has {
		return -1, errors.New("name is being used")
	}

	var id int
	insertVisitorQuery := `INSERT INTO Visitors (name, email, password) values ($1, $2, $3) RETURNING id`

	row = r.db.QueryRow(insertVisitorQuery, visitor.Name, visitor.Email, visitor.Password)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *AuthPostgres) GetVisitor(name, password string) (entity.Visitor, error) {
	var user entity.Visitor
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND password=$2", visitorsTable)
	err := r.db.Get(&user, query, name, password)
	if err != nil {
		return entity.Visitor{}, err
	}
	return user, nil
}

func (r *AuthPostgres) CreateCoach(coach entity.Coach) (int, error) {
	var has bool = false
	visitorExistsQuery := `select exists(select * from Visitors where name = $1) or
						exists(select * from Coaches where name = $1)`

	row := r.db.QueryRow(visitorExistsQuery, coach.Name)
	if err := row.Scan(&has); err != nil {
		return -1, err
	}

	if has {
		return -1, errors.New("name is being used")
	}

	var id int
	insertCoachQuery := `INSERT INTO Coaches (name, email, password) values ($1, $2, $3) RETURNING id`

	row = r.db.QueryRow(insertCoachQuery, coach.Name, coach.Email, coach.Password)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *AuthPostgres) GetCoach(name, password string) (entity.Coach, error) {
	var user entity.Coach
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND password=$2", coachesTable)
	err := r.db.Get(&user, query, name, password)
	if err != nil {
		return entity.Coach{}, err
	}
	return user, nil
}
