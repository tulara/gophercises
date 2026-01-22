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
	NextCursor string `json:"next_cursor,omitempty"`
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

	startFrom := 1 // Cafes IDs start from 1 not 0
	if cursor != "" {
		startFrom, err = strconv.Atoi(cursor)
		if err != nil {
			http.Error(w, "Unable to parse page size as integer", http.StatusBadRequest)
			return
		}
	}

	// Fetch an additional cafe so we know what the next cursor will be
	cafes := h.store.GetCafes(size+1, startFrom)

	response := &GetCafesDTO{
		Data: cafes,
	}

	if len(cafes) == size+1 {
		response = &GetCafesDTO{
			Data: cafes[0:size],
			Pagination: Pagination{
				NextCursor: strconv.Itoa(cafes[size].ID),
			},
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
