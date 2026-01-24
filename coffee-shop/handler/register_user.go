package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tulara/coffeeshop/auth"
	"github.com/tulara/coffeeshop/domain"
)

type RegisterUserResponse struct {
	Token string `json:"token"`
}

func (h *Handler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid request body, username and password are required.", http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		// Worth validating assumption that this won't send passwords to logs... or obfusicating logs
		fmt.Printf("Failed to hash password: %v", err)
		http.Error(w, "Failed Auth", http.StatusInternalServerError)
	}
	h.store.CreateUser(user.Username, hashedPassword)

	token, err := auth.CreateToken(user.Username)
	if err != nil {
		fmt.Printf("Failed to create JWT: %v", err)
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	response := RegisterUserResponse{
		Token: token,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
