package service

import (
	"context"
	"ecom-go/internal/models"
	"time"
)

type Storage interface {
	Create(task *models.ToDoRequest) *models.ToDo
	GetAll() []*models.ToDo
	GetByID(id int) (*models.ToDo, error)
	Update(id int, task *models.ToDoRequest) error
	Delete(id int) error
}

type ToDoService struct {
	storage Storage
}

func NewService(storage Storage) *ToDoService {
	return &ToDoService{storage: storage}
}

func (s *ToDoService) Create(ctx context.Context, task *models.ToDoRequest) (*models.ToDo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	done := make(chan *models.ToDo, 1)
	go func() {
		done <- s.storage.Create(task)
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case t := <-done:
		return t, nil
	}
}

func (s *ToDoService) GetAll(ctx context.Context) ([]*models.ToDo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	done := make(chan []*models.ToDo, 1)
	go func() {
		done <- s.storage.GetAll()
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case tasks := <-done:
		return tasks, nil
	}
}

func (s *ToDoService) GetByID(ctx context.Context, id int) (*models.ToDo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	done := make(chan *models.ToDo, 1)
	errCh := make(chan error, 1)
	go func() {
		task, err := s.storage.GetByID(id)
		if err != nil {
			errCh <- err
			return
		}
		done <- task
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errCh:
		return nil, err
	case task := <-done:
		return task, nil
	}
}

func (s *ToDoService) Update(ctx context.Context, id int, task *models.ToDoRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- s.storage.Update(id, task)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}

func (s *ToDoService) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- s.storage.Delete(id)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}
