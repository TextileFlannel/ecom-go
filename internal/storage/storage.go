package storage

import (
	"ecom-go/internal/models"
	"errors"
	"sync"
)

type MemoryStorage struct {
	data map[int]*models.ToDo
	id   int
	mu   sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]*models.ToDo),
		id:   1,
	}
}

func (s *MemoryStorage) Create(task *models.ToDoRequest) error {
	if task.Title == "" {
		return errors.New("Title is required")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[s.id] = &models.ToDo{
		ID:    s.id,
		Title: task.Title,
		Body:  task.Body,
	}
	s.id++

	return nil
}

func (s *MemoryStorage) GetAll() []*models.ToDo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]*models.ToDo, 0, len(s.data))
	for _, task := range s.data {
		tasks = append(tasks, task)
	}

	return tasks
}

func (s *MemoryStorage) GetByID(id int) (*models.ToDo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.data[id]
	if !ok {
		return nil, errors.New("Not found")
	}
	return task, nil
}

func (s *MemoryStorage) Update(id int, task *models.ToDoRequest) error {
	if task.Title == "" {
		return errors.New("Title is required")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[id]; !ok {
		return errors.New("Not found")
	}

	s.data[id].Title = task.Title
	s.data[id].Body = task.Body

	return nil
}

func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[id]; !ok {
		return errors.New("Not found")
	}

	delete(s.data, id)

	return nil
}
