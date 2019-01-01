package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/askft/wloggr/api/store"
	"github.com/askft/wloggr/api/util"

	"github.com/go-chi/chi"
)

// WorkoutRoutes ...
func WorkoutRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(VerifyToken)
	r.Post("/", createWorkout)
	r.Put("/{date}", updateWorkout)
	r.Put("/{date}/new", updateWorkoutDate)
	r.Delete("/{date}", deleteWorkout)
	r.Get("/{date}", getWorkout)
	r.Get("/", getWorkouts)
	r.Get("/dates", getWorkoutDates)
	return r
}

func createWorkout(w http.ResponseWriter, r *http.Request) {
	userID := ctxString(r, ckUserID)
	date := time.Now().Format(util.SQLDateTime)
	fmt.Println(date)

	b, err := ioutil.ReadAll(r.Body) // TODO this looks dangerous
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = store.Store.CreateWorkout(userID, date, string(b[:])); err != nil {
		fmt.Println(err)
		return
	}

	// TODO don't send a god damn plain string
	util.SendAsJSON(w, date)
}

func updateWorkout(w http.ResponseWriter, r *http.Request) {
	userID := ctxString(r, ckUserID)
	date := chi.URLParam(r, "date")
	workout, err := ioutil.ReadAll(r.Body) // TODO this looks dangerous
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("\nuserID\n%s\n\nDate\n%s\n\nWorkout\n%s\n\n", userID, date, workout)
	if err = store.Store.UpdateWorkout(userID, date, string(workout[:])); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateWorkoutDate(w http.ResponseWriter, r *http.Request) {
	userID := ctxString(r, ckUserID)
	date := chi.URLParam(r, "date")
	s := struct {
		NewDate string `json:"newDate"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.Store.UpdateWorkoutDate(userID, date, s.NewDate)
}

func deleteWorkout(w http.ResponseWriter, r *http.Request) {
	userID := ctxString(r, ckUserID)
	date := chi.URLParam(r, "date")
	// TODO Parse this into a date object and return error if fail
	fmt.Println("Got date [" + date + "].")
	if err := store.Store.DeleteWorkout(userID, date); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendAsJSON(w, fmt.Sprintf(
		"Deleted workout with userID(%s) and date(%s).",
		userID, date))
}

func getWorkout(w http.ResponseWriter, r *http.Request) {
	userID := ctxString(r, ckUserID)
	date := chi.URLParam(r, "date")
	fmt.Println("Got date [" + date + "].")
	workout, err := store.Store.GetWorkout(userID, date)
	if err != nil {
		// TODO use 404 instead?
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// s := models.Workout{}
	// json.Unmarshal([]byte(workout), &s)
	// util.SendAsJSON(w, s)
	// fmt.Printf("\n%v\n", workout)
	w.Write([]byte(workout))
}

func getWorkouts(w http.ResponseWriter, r *http.Request) {
	userID := ctxString(r, ckUserID)
	workouts, err := store.Store.GetWorkouts(userID)
	if err != nil {
		// TODO use 404 instead?
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendAsJSON(w, workouts)
}

func getWorkoutDates(w http.ResponseWriter, r *http.Request) {
	userID := ctxString(r, ckUserID)
	dates, err := store.Store.GetWorkoutDates(userID)
	if err != nil {
		// TODO use 404 instead?
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s := struct {
		WorkoutDates []string `json:"workoutDates"`
	}{dates}
	util.SendAsJSON(w, s)
}
