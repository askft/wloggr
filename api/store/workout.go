package store

import (
	"encoding/json"
	"fmt"

	"github.com/askft/wloggr/api/models"
)

// CreateWorkout ...
func (p *DB) CreateWorkout(userID, date, jsn string) error {
	q := `INSERT INTO workouts (userId, date, jdoc) VALUES (?, ?, ?)`
	_, err := p.db.Exec(q, userID, date, jsn)
	return err
}

// UpdateWorkout ...
func (p *DB) UpdateWorkout(userID, date, jsn string) error {
	q := `UPDATE workouts SET jdoc = (?) WHERE userId = (?) AND date = (?)`
	_, err := p.db.Exec(q, jsn, userID, date)
	return err
}

// UpdateWorkoutDate ...
func (p *DB) UpdateWorkoutDate(userID, date, newDate string) error {
	q := `UPDATE workouts SET date = (?) WHERE userId = (?) AND date = (?)`
	_, err := p.db.Exec(q, newDate, userID, date)
	return err
}

// DeleteWorkout ...
func (p *DB) DeleteWorkout(userID, date string) error {
	q := `DELETE FROM workouts WHERE userID = (?) AND date = (?);`
	_, err := p.db.Exec(q, userID, date)
	return err
}

// GetWorkout ...
func (p *DB) GetWorkout(userID, date string) (string, error) {
	q := `SELECT jdoc FROM workouts WHERE userId = (?) AND date = (?)`
	var w string
	err := p.db.Get(&w, q, userID, date)
	if err != nil {
		return w, err
	}
	return w, nil
}

// GetWorkouts ...
func (p *DB) GetWorkouts(userID string) ([]models.Workout, error) {
	q := `SELECT date, jdoc FROM workouts WHERE userId = (?)`
	rows, err := p.db.Query(q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 	Initializing the list with an empty Workout struct
	// is useful when Marshaling to JSON.
	// 	ACTUALLY – this might be completely useless,
	// as it is the list that gets Marshalled... I'll see.
	workouts := []models.Workout{models.CreateEmptyWorkout()}
	for rows.Next() {
		var date, jsn string
		if err := rows.Scan(&date, &jsn); err != nil {
			return nil, err
		}
		partialWorkout := struct {
			Exercises []models.Exercise `json:"exercises"`
		}{}
		json.Unmarshal([]byte(jsn), &partialWorkout)
		fmt.Println(partialWorkout)
		workout := models.Workout{
			Date:      date,
			Exercises: partialWorkout.Exercises,
		}
		workout.Date = date
		blah, _ := json.MarshalIndent(workout, "", "  ")
		fmt.Printf("\n%s\n", blah)
		workouts = append(workouts, workout)
	}
	b, _ := json.MarshalIndent(workouts, "", "  ")
	fmt.Println(string(b))
	return workouts, nil
}

// GetWorkoutDates ...
func (p *DB) GetWorkoutDates(userID string) ([]string, error) {
	q := `SELECT date FROM workouts WHERE userId = (?)`
	rows, err := p.db.Query(q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dates := []string{}
	for rows.Next() {
		var date string
		if err := rows.Scan(&date); err != nil {
			return nil, err
		}
		dates = append(dates, date)
	}
	return dates, nil
}
