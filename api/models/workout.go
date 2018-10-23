package models

type Workout struct {
	Date      string     `json:"date"`
	Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
	Name string `json:"name"`
	Sets []Set  `json:"sets"`
}

type Set struct {
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
}

// CreateEmptyWorkout (DEPRECATED) returns a workout object with
// each array the length of one and each leaf set to the zero value.
func CreateEmptyWorkout() Workout {
	return Workout{
		Date: "",
		Exercises: []Exercise{
			Exercise{
				Name: "",
				Sets: []Set{
					Set{
						Reps:   0,
						Weight: 0,
					},
				},
			},
		},
	}
}
