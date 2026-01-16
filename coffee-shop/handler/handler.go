package handler

import "github.com/tulara/coffeeshop/store"

type Handler struct {
	store store.Store
}

func NewHandler(store store.Store) *Handler {
	return &Handler{
		store: store,
	}
}
