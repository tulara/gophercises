package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tulara/coffeeshop/domain"
)

func (h *Handler) HandleCreateCafe(w http.ResponseWriter, r *http.Request) {
	// use json decoder (rather than unmarshal) when json data is incoming from a stream
	// and does not need to be fully loaded into memory at once.
	var cafe domain.Cafe
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cafe)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if cafe.ID == "" {
		http.Error(w, "Cafe ID is required", http.StatusBadRequest)
		return
	}

	// Check if cafe already exists to determine status code
	existingCafe := h.store.GetCafe(cafe.ID)
	h.store.CreateCafe(&cafe)

	if existingCafe != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Cafe updated: %+v\n", cafe)
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Cafe created successfully: %+v\n", cafe)
	}
}
