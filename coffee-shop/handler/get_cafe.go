package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) HandleGetCafes(w http.ResponseWriter, r *http.Request) {
	cafes := h.store.GetCafes()

	cafesJson, err := json.Marshal(cafes)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cafesJson)
}

func (h *Handler) HandleGetCafe(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cafe := h.store.GetCafe(id)
	if cafe == nil {
		http.Error(w, "Cafe not found", http.StatusNotFound)
		return
	}

	cafeJson, err := json.Marshal(cafe)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cafeJson)
}
