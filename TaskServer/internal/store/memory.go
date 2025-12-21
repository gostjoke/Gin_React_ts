package store

import (
	"errors"
	"sync"
	"time"

	"TaskServer/internal/model"
)

var ErrNotFound = errors.New("not found")

type MemoryStore struct {
	mu    sync.RWMutex
	tasks map[string]model.Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{tasks: make(map[string]model.Task)}
}

func (s *MemoryStore) Create(t model.Task) (model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	s.tasks[t.ID] = t
	return t, nil
}

func (s *MemoryStore) Get(id string) (model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tasks[id]
	if !ok {
		return model.Task{}, ErrNotFound
	}
	return t, nil
}

func (s *MemoryStore) List() ([]model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]model.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	return out, nil
}

func (s *MemoryStore) Update(t model.Task) (model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[t.ID]; !ok {
		return model.Task{}, ErrNotFound
	}
	s.tasks[t.ID] = t
	return t, nil
}
