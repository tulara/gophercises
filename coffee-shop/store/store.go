package store

import "github.com/tulara/coffeeshop/domain"

type Store interface {
	CreateCafe(cafe *domain.Cafe)
	GetCafe(id string) *domain.Cafe

	// Supports cursor based pagination.
	// Returns size number of cafes starting from the cursor specified,
	// plus an extra cafe if there are more available.
	GetCafes(size int, cursor string) []*domain.Cafe
}
