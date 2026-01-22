package store

import (
	"sync"

	"github.com/tulara/coffeeshop/domain"
)

type MemoryStore struct {
	// RWMutex allows arbitrary reads alongside a writer thread.
	// Really its probably simpler to start with a mutex, as there's a small performance overhead
	// and some likelihood of mixing up read and write locks.
	// But this is a good way to refresh my memory.
	mut   sync.RWMutex
	cafes map[int]*domain.Cafe
}

func NewMemoryStore() Store {
	return &MemoryStore{
		cafes: map[int]*domain.Cafe{},
	}
}

func (s *MemoryStore) CreateCafe(cafe *domain.Cafe) {
	s.mut.Lock() //exclusive (write) lock.
	defer s.mut.Unlock()
	s.cafes[cafe.ID] = cafe
}

func (s *MemoryStore) GetCafe(id int) *domain.Cafe {
	s.mut.RLock()
	defer s.mut.RUnlock()
	return s.cafes[id]
}

// Assume, make sure, s.cafes is ordered by ID.
func (s *MemoryStore) GetCafes(size int, startFrom int) []*domain.Cafe {
	s.mut.RLock()
	defer s.mut.RUnlock()

	cafes := []*domain.Cafe{}

	// 5, 1 | 3
	// 2, 1 | 3
	// 1, 3 | 4

	// 2, 3 | 6
	// 3, 4 | 5
	// 2, 3 | 3

	end := startFrom + size
	if startFrom+size > len(s.cafes) {
		end = len(s.cafes) + 1
	}

	for i := startFrom; i < end; i++ {
		cafes = append(cafes, s.cafes[i])
	}
	return cafes
}
