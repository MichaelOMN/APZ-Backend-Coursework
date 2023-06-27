package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	activitiesTable     = "activities"
	activitystatesTable = "activitystates"
	activityusageTable  = "activityusage"
	attendanceTable     = "attendance"
	clubsTable          = "clubs"
	coachesTable        = "coaches"
	physicalinfoTable   = "physicalinfo"
	physicalstatesTable = "physicalstates"
	statestypesTable    = "statestypes"
	trainingsTable      = "trainings"
	visitorsTable       = "visitors"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// var config Config = Config{
// 	Host:     viper.GetString("db.host"),
// 	Port:     viper.GetString("db.port"),
// 	Username: viper.GetString("db.username"),
// 	DBName:   viper.GetString("db.dbname"),
// 	SSLMode:  viper.GetString("db.sslmode"),
// 	Password: os.Getenv("DB_PASSWORD"),
// }

// func GetDB() (*sqlx.DB, error) {
// 	db, err := NewPostgresDB(config)

// 	if err != nil {
// 		logrus.Fatalf("failed to initialize db: %s", err.Error())
// 		return nil, err
// 	}

// 	return db, nil
// }
