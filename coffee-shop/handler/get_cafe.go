package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tulara/coffeeshop/domain"
)

type GetCafesDTO struct {
	Data       []*domain.Cafe `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

type Pagination struct {
	NextCursor string `json:"next_cursor"`
}

// Page size defaults to 5.
func (h *Handler) HandleGetCafes(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	pageSize := queryParams.Get("page_size")
	cursor := queryParams.Get("cursor")

	size := 5
	var err error

	if pageSize != "" {
		size, err = strconv.Atoi(pageSize)
		if err != nil {
			http.Error(w, "Unable to parse page size as integer", http.StatusBadRequest)
			return
		}
	}

	cafes := h.store.GetCafes(size, cursor)

	response := GetCafesDTO{
		Data: cafes,
	}
	if len(cafes) > size {
		response.Pagination = Pagination{
			NextCursor: cafes[size].ID,
		}
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
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
