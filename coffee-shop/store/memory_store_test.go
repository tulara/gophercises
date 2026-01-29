package store

import (
	"testing"

	"github.com/tulara/coffeeshop/domain"
)

func TestDefaultPageSizeAndCursor(t *testing.T) {
	cafes := []domain.Cafe{
		{Name: "Blue Bottle", ID: 1},
		{Name: "Intelligentsia", ID: 2},
		{Name: "Counter Culture", ID: 3},
	}

	store := NewMemoryStore()
	for _, c := range cafes {
		store.UpsertCafe(&c)
	}

	retrievedCafes := store.GetCafes(5, 1)
	if len(retrievedCafes) != 3 {
		t.Errorf("Expected 3 cafes but there were %d", len(retrievedCafes))
	}
}

func TestSmallerPageSizeAndCursor(t *testing.T) {
	cafes := []domain.Cafe{
		{Name: "Blue Bottle", ID: 1},
		{Name: "Intelligentsia", ID: 2},
		{Name: "Counter Culture", ID: 3},
	}

	store := NewMemoryStore()
	for _, c := range cafes {
		store.UpsertCafe(&c)
	}

	retrievedCafes := store.GetCafes(2, 1)
	if len(retrievedCafes) != 2 {
		t.Errorf("Expected 2 cafes but there were %d", len(retrievedCafes))
	}
}
