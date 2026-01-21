package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) HandleGetCafe(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "An id is required to fetch a cafe", http.StatusBadRequest)
		return
	}

	cafeId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Unable to parse path variable as an integer", http.StatusBadRequest)
		return
	}

	cafe := h.store.GetCafe(cafeId)
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
