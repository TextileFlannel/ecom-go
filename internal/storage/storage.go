package storage

import (
	"ecom-go/internal/models"
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
