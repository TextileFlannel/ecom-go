package service

import (
	"context"
	"ecom-go/internal/models"
	"ecom-go/internal/storage"
	"time"
)

type ToDoService struct {
	storge *storage.MemoryStorage
}

func NewService(storge *storage.MemoryStorage) *ToDoService {
	return &ToDoService{storge: storge}
}

func (s *ToDoService) Create(ctx context.Context, task *models.ToDoRequest) (*models.ToDo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	done := make(chan *models.ToDo, 1)
	go func() {
		done <- s.storge.Create(task)
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
		done <- s.storge.GetAll()
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
		task, err := s.storge.GetByID(id)
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
		done <- s.storge.Update(id, task)
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
		done <- s.storge.Delete(id)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}
