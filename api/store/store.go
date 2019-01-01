package store

import (
	_ "github.com/go-sql-driver/mysql" // SQLx uses this package.
	"github.com/jmoiron/sqlx"

	"github.com/askft/wloggr/api/models"
)

// Store is a package-level reference to a database handle.
// Use SetupDB or SetupMockDB to initialize.
var Store store

type store interface {
	NewUser(*models.User) error
	UpdateUserFullName(userID, fullName string) error
	GetUserByEmail(string) (*models.User, error)
	GetUserByID(string) (*models.User, error)
	GetUsers() ([]*models.User, error)

	CreateWorkout(userID, date, jsn string) error
	UpdateWorkout(userID, date, jsn string) error
	UpdateWorkoutDate(userID, date, newDate string) error
	DeleteWorkout(userID, date string) error
	GetWorkout(userID, date string) (string, error)
	GetWorkouts(userID string) ([]models.Workout, error)
	GetWorkoutDates(userID string) ([]string, error)
}

// DB is a wrapper for the database connection.
type DB struct {
	db *sqlx.DB
}

// SetupDB sets up a new database connection. A wrapper around the
// connection is accessible via the package-level variable `Store`.
func SetupDB(dataSourceName string) error {
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return err
	}
	db.MapperFunc(func(s string) string { return s })
	db.MustExec(schemaWorkouts)
	db.MustExec(schemaUsers)
	Store = &DB{db}
	return nil
}

var schemaWorkouts = `
CREATE TABLE IF NOT EXISTS workouts (
	id				INT 			NOT NULL PRIMARY KEY AUTO_INCREMENT,
	userId			VARCHAR(255)	NOT NULL,
	date			DATETIME	 	NOT NULL,
	jdoc			JSON			NOT NULL
);`

var schemaUsers = `
CREATE TABLE IF NOT EXISTS users (
	id				INT				NOT NULL PRIMARY KEY AUTO_INCREMENT,
	userId			VARCHAR(255)	NOT NULL UNIQUE,
	email			VARCHAR(255)	NOT NULL UNIQUE,
	hash			VARCHAR(255)	NOT NULL,
	fullName		VARCHAR(255)	NOT NULL
);`
