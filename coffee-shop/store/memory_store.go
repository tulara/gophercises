package store

import (
	"sync"

	"github.com/tulara/coffeeshop/domain"
)

type MemoryStore struct {
	// RWMutex allows arbitrary reads alongside a writer thread.
	// Really its probably simpler to start with a mutex, as there's a small performance overhead
	// and some likelyhood of mixing up read and write locks.
	// But this is a good way to refresh my memory.
	mut   sync.RWMutex
	cafes map[string]*domain.Cafe
}

func NewMemoryStore() Store {
	return &MemoryStore{
		cafes: map[string]*domain.Cafe{},
	}
}

func (s *MemoryStore) CreateCafe(cafe *domain.Cafe) {
	s.mut.Lock() //exclusive (write) lock.
	defer s.mut.Unlock()
	s.cafes[cafe.ID] = cafe
}

func (s *MemoryStore) GetCafe(id string) *domain.Cafe {
	s.mut.RLock()
	defer s.mut.RUnlock()
	return s.cafes[id]
}
