package handlers

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	//parse json body!
	var req CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	//validation
	if req.Email == "" || req.Password == "" {
		http.Error(w, "email and password required", http.StatusBadRequest)
		return
	}

	//respomse
	response := map[string]string{
		"message": "user created",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
