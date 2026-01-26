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
	cafeMut sync.RWMutex
	userMut sync.RWMutex
	cafes   map[int]*domain.Cafe
	users   map[string]string
}

func NewMemoryStore() Store {
	return &MemoryStore{
		cafes: map[int]*domain.Cafe{},
		users: map[string]string{},
	}
}

func (s *MemoryStore) CreateCafe(cafe *domain.Cafe) {
	s.cafeMut.Lock() //exclusive (write) lock.
	defer s.cafeMut.Unlock()
	s.cafes[cafe.ID] = cafe
}

func (s *MemoryStore) GetCafe(id int) *domain.Cafe {
	s.cafeMut.RLock()
	defer s.cafeMut.RUnlock()
	return s.cafes[id]
}

func (s *MemoryStore) GetCafes(size int, startFrom int) []*domain.Cafe {
	s.cafeMut.RLock()
	defer s.cafeMut.RUnlock()

	cafes := []*domain.Cafe{}

	end := startFrom + size
	if startFrom+size > len(s.cafes) {
		end = len(s.cafes) + 1
	}

	for i := startFrom; i < end; i++ {
		cafes = append(cafes, s.cafes[i])
	}
	return cafes
}

func (s *MemoryStore) CreateUser(username string, password string) {
	s.userMut.Lock()
	defer s.userMut.Unlock()

	s.users[username] = password
}

func (s *MemoryStore) GetUser(username string) *domain.User {
	s.userMut.RLock()
	defer s.userMut.RUnlock()

	password, ok := s.users[username]
	if !ok {
		return nil
	}

	return &domain.User{
		Username: username,
		Password: password,
	}
}
