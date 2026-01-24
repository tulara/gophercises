package store

import "github.com/tulara/coffeeshop/domain"

type Store interface {
	CreateCafe(cafe *domain.Cafe)
	GetCafe(id int) *domain.Cafe

	// Supports cursor based pagination.
	// Returns size number of cafes starting from the cursor specified,
	// plus an extra cafe if there are more available.
	GetCafes(size int, cursor int) []*domain.Cafe

	CreateUser(username string, passwword string)
}
