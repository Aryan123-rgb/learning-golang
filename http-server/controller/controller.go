package controller

import (
	"encoding/json"
	"fmt"
	"http-server/db"
	"net/http"
	"time"
)

// This file contains handler function for different routes

type Controller struct {
	Db *db.Database
}

func NewController(database *db.Database) *Controller {
	return &Controller{
		Db: database,
	}
}

// GET -> /
// Returns a simple welcome message
func (h *Controller) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintf(w, "Welcome to our custom http server")
}

// GET -> /time
// Returns the current time in formatted json object
func (h *Controller) GetTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)
	json.NewEncoder(w).Encode(map[string]string{"time": currentTime})
}

// GET -> /users
// List all the users in the map
func (h *Controller) UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	users := h.Db.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

// POST -> /user
// Recieve a json object user in req body and add it to the db
func (h *Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid Request body %v", err.Error()), http.StatusBadRequest)
	}

	user := h.Db.CreateNewUser(req.Email, req.Username)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}
