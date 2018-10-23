package route

import (
	"encoding/json"
	"net/http"

	"wloggr/api/models"
	"wloggr/api/services"
	"wloggr/api/store"
	"wloggr/api/util"

	"github.com/go-chi/chi"
)

// UserRoutes ...
func UserRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/signup", signup)
	r.Post("/signin", signin)
	r.Get("/users", getUsers)
	r.With(VerifyToken).Put("/profile/fullname", updateUser)
	r.With(VerifyToken).Get("/profile", profile)
	return r
}

func profile(w http.ResponseWriter, r *http.Request) {
	userID := ctxString(r, ckUserID)
	user, err := store.Store.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendAsJSON(w, user)
}

func signup(w http.ResponseWriter, r *http.Request) {
	c := &models.Signup{}
	if err := json.NewDecoder(r.Body).Decode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := services.Signup(c)
	if err != nil {
		http.Error(w, err.Msg, err.Code)
		return
	}
}

func signin(w http.ResponseWriter, r *http.Request) {
	c := &models.Signin{}
	if err := json.NewDecoder(r.Body).Decode(c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := services.Signin(c)
	if err != nil {
		http.Error(w, err.Msg, err.Code)
		return
	}

	signature, err := services.GetToken(u.UserID)
	if err != nil {
		http.Error(w, err.Msg, err.Code)
		return
	}

	util.SendAsJSON(w, signature)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	s := struct {
		FullName string `json:"fullName"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userID := ctxString(r, ckUserID)
	if err := store.Store.UpdateUserFullName(userID, s.FullName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := store.Store.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendAsJSON(w, users)
}
