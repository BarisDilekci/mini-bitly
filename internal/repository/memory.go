// Package repository defines the storage interface and implementations.
package repository

import (
	"mini-bitly/internal/domain"
	"sync"
)

// Repository defines the contract for storing and retrieving Links.
// It allows the storage backend to be swapped (e.g., InMemory, Postgres, Redis)
// without changing the business logic.
type Repository interface {
	// Save persists a Link to the storage.
	Save(link *domain.Link) error

	// FindByCode retrieves a Link by its short code.
	// Returns domain.ErrNotFound if the code represents no link.
	FindByCode(code string) (*domain.Link, error)
}

// memoryRepository is a thread-safe, in-memory implementation of the Repository interface.
// It uses a map to store data and RWMutex for concurrency control.
// This is suitable for testing or small workloads where data persistence across restarts is not required.
type memoryRepository struct {
	mu    sync.RWMutex
	links map[string]*domain.Link
}

// NewMemoryRepository creates a new instance of an in-memory repository.
func NewMemoryRepository() Repository {
	return &memoryRepository{
		links: make(map[string]*domain.Link),
	}
}

// Save stores the link in memory.
// It acquires a write lock to ensure thread safety.
func (r *memoryRepository) Save(link *domain.Link) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.links[link.Code] = link
	return nil
}

// FindByCode retrieves the link from memory.
// It acquires a read lock to allow multiple concurrent readers.
func (r *memoryRepository) FindByCode(code string) (*domain.Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	link, ok := r.links[code]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return link, nil
}
