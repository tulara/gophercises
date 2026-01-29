package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tulara/coffeeshop/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Invalid request body, username and password are required.", http.StatusBadRequest)
		return
	}

	user := h.store.GetUser(req.Username)
	if user == nil {
		// Shouldn't tell user whether username exists or not.
		// Ideally capture email and if they don't get an email they can go through
		// a reset flow.
		http.Error(w, "Failed Auth", http.StatusUnauthorized)
		return
	}

	isMatch, err := auth.VerifyPassword(user.Password, req.Password)
	if err != nil {
		fmt.Printf("Error verifying password: %v", err)
		http.Error(w, "Failed Auth", http.StatusUnauthorized)
		return
	}
	if !isMatch {
		http.Error(w, "Failed Auth", http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateToken(user.Username)
	if err != nil {
		fmt.Printf("Failed to create JWT: %v", err)
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
