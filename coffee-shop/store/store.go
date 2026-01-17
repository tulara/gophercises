package store

import "github.com/tulara/coffeeshop/domain"

type Store interface {
	CreateCafe(cafe *domain.Cafe)
	GetCafe(id string) *domain.Cafe
	GetCafes() []*domain.Cafe
}
