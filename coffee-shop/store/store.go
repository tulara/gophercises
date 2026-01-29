package store

import "github.com/tulara/coffeeshop/domain"

type Store interface {
	// UpsertCafe returns true if user already existed.
	UpsertCafe(cafe *domain.Cafe) bool
	GetCafe(id int) *domain.Cafe

	// GetCafes supports cursor based pagination.
	// Returns size number of cafes starting from the cursor specified,
	// plus an extra cafe if there are more available.
	GetCafes(size int, cursor int) []*domain.Cafe

	// UpsertUser returns true if user already existed.
	CreateUser(username string, password string)
	GetUser(username string) *domain.User
}
