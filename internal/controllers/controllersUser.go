package controllers

import (
	"blog/internal/models"
	"blog/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

type UserController struct {
	repo *repository.UserRepository
}

func NewUserController(repo *repository.UserRepository) *UserController {
	return &UserController{repo}
}

func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := uc.repo.CreateUser(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Get all users from the repository
	users, err := uc.repo.GetAllUsers(r.Context())
	if err != nil {
		// Handle error and return a JSON response with a meaningful error message
		uc.handleError(w, err, http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the users to JSON and write to the response writer
	if err := json.NewEncoder(w).Encode(users); err != nil {
		// Handle encoding error
		uc.handleError(w, err, http.StatusInternalServerError)
		return
	}

	// Log the users for debugging purposes (optional)
	log.Printf("Fetched users: %+v\n", users)
}

// Helper function to handle errors and write a JSON response
func (uc *UserController) handleError(w http.ResponseWriter, err error, statusCode int) {
	// Create a JSON error response
	errorResponse := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	// Set the HTTP status code
	w.WriteHeader(statusCode)

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the error response to JSON and write to the response writer
	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		// Log the encoding error
		log.Printf("Error encoding error response: %v\n", err)
	}
}
