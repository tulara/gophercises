package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tulara/coffeeshop/auth"
	"github.com/tulara/coffeeshop/domain"
)

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

	w.WriteHeader(http.StatusCreated)
}
