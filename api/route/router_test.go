package route

import (
	"bytes"
	"fmt"
	"github.com/askft/wloggr/api/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	store.SetupMockDB()
	fmt.Printf("%+v\n", store.Store)
}

func TestCreateWorkout(t *testing.T) {
	buf := bytes.NewBufferString(store.RouterTestWorkout())
	req, err := http.NewRequest("POST", "", buf)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(createWorkout)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetWorkouts(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(getWorkouts)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// expected := `{go get git}`

	fmt.Println(rr.Body.String())
	return

	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}
