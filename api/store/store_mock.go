package store

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"wloggr/api/models"
)

// MockDB ...
type MockDB struct{}

// SetupMockDB ...
func SetupMockDB() {
	Store = &MockDB{}
	fmt.Println("Finished setting up mock database.")
}

// NewUser ...
func (p *MockDB) NewUser(u *models.User) error {
	fmt.Println("MockStore: NewUser")
	return nil
}

// UpdateUserFullName ...
func (p *MockDB) UpdateUserFullName(userID, fullName string) error {
	fmt.Println("MockStore: UpdateUserFullName")
	return nil
}

// GetUserByID ...
func (p *MockDB) GetUserByID(userID string) (*models.User, error) {
	fmt.Println("MockStore: GetUserByID", userID)
	u := &models.User{}
	return u, nil
}

// GetUserByEmail ...
func (p *MockDB) GetUserByEmail(email string) (*models.User, error) {
	fmt.Println("MockStore: GetUserByEmail", email)
	u := &models.User{}
	return u, nil
}

// GetUsers ...
func (p *MockDB) GetUsers() ([]*models.User, error) {
	fmt.Println("MockStore: GetUsers")
	us := []*models.User{}
	return us, nil
}

//
// Workouts -------------------------------------

// RouterTestWorkout ...
func RouterTestWorkout() string {
	return removeWhitespace(`{
		"exercises": [
			{
				"name": "ohp",
				"sets": [
					{ "reps": 5, "weight": 100 },
					{ "reps": 3, "weight": 110 }
				]
			}
		]
	}`)
}
func removeWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// CreateWorkout ...
func (p *MockDB) CreateWorkout(userID, date, jsn string) error {
	fmt.Println("MockStore: CreateWorkout with", userID, date, jsn)
	if jsn != RouterTestWorkout() {
		return errors.New("invalid")
	}
	return nil
}

// UpdateWorkout ...
func (p *MockDB) UpdateWorkout(userID, date, jsn string) error {
	fmt.Printf("MockStore: UpdateWorkout id(%s), date(%s) with %s\n",
		userID, date, jsn)
	return nil
}

// UpdateWorkoutDate ...
func (p *MockDB) UpdateWorkoutDate(userID, date, newDate string) error {
	fmt.Printf("MockStore: UpdateWorkoutDate userID(%s), date(%s), newDate(%s)\n",
		userID, date, newDate)
	return nil
}

// DeleteWorkout ...
func (p *MockDB) DeleteWorkout(userID, date string) error {
	fmt.Printf("MockStore: DeleteWorkout userID(%s), date(%s)\n",
		userID, date)
	return nil
}

// GetWorkout ...
func (p *MockDB) GetWorkout(userID, date string) (string, error) {
	fmt.Printf("MockStore: GetWorkout userID(%s), date(%s)\n",
		userID, date)
	return "", nil
}

// GetWorkouts ...
func (p *MockDB) GetWorkouts(userID string) ([]models.Workout, error) {
	fmt.Printf("MockStore: GetWorkouts userID(%s)\n", userID)
	workouts := []models.Workout{}
	return workouts, nil
}

// GetWorkoutDates ...
func (p *MockDB) GetWorkoutDates(userID string) ([]string, error) {
	fmt.Printf("MockStore: GetWorkoutDates userID(%s)\n", userID)
	dates := []string{}
	return dates, nil
}
