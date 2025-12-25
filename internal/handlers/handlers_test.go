package handlers

import (
	"bytes"
	"context"
	"ecom-go/internal/models"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
)

type mockToDoService struct {
	createFunc  func(ctx context.Context, task *models.ToDoRequest) (*models.ToDo, error)
	getAllFunc  func(ctx context.Context) ([]models.ToDo, error)
	getByIDFunc func(ctx context.Context, id int) (*models.ToDo, error)
	updateFunc  func(ctx context.Context, id int, task *models.ToDoRequest) error
	deleteFunc  func(ctx context.Context, id int) error
}

func (m *mockToDoService) Create(ctx context.Context, task *models.ToDoRequest) (*models.ToDo, error) {
	return m.createFunc(ctx, task)
}
func (m *mockToDoService) GetAll(ctx context.Context) ([]models.ToDo, error) {
	return m.getAllFunc(ctx)
}
func (m *mockToDoService) GetByID(ctx context.Context, id int) (*models.ToDo, error) {
	return m.getByIDFunc(ctx, id)
}
func (m *mockToDoService) Update(ctx context.Context, id int, task *models.ToDoRequest) error {
	return m.updateFunc(ctx, id, task)
}
func (m *mockToDoService) Delete(ctx context.Context, id int) error {
	return m.deleteFunc(ctx, id)
}

func TestCreateHandler(t *testing.T) {
	tests := []struct {
		name       string
		body       interface{}
		setupMock  func(*mockToDoService)
		wantStatus int
		wantBody   string
	}{
		{
			name: "success",
			body: models.ToDoRequest{Title: "Test", Body: "Desc"},
			setupMock: func(m *mockToDoService) {
				m.createFunc = func(ctx context.Context, task *models.ToDoRequest) (*models.ToDo, error) {
					return &models.ToDo{ID: 1, Title: "Test", Body: "Desc", IsDone: false}, nil
				}
			},
			wantStatus: 201,
			wantBody:   `"title":"Test"`,
		},
		{
			name:       "empty title validation",
			body:       models.ToDoRequest{Title: "", Body: "Desc"},
			setupMock:  func(m *mockToDoService) {},
			wantStatus: 400,
			wantBody:   "title is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockToDoService{}
			tt.setupMock(mock)
			handler := &ToDoHandlers{service: mock}

			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
			rr := httptest.NewRecorder()

			handler.Create(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rr.Code, tt.wantStatus)
			}
			if !bytes.Contains(rr.Body.Bytes(), []byte(tt.wantBody)) {
				t.Errorf("body = %q, want %q", rr.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestUpdateHandler(t *testing.T) {
	mock := &mockToDoService{}
	mock.updateFunc = func(ctx context.Context, id int, task *models.ToDoRequest) error {
		if task.Title == "" {
			return errors.New("not found")
		}
		return nil
	}

	handler := &ToDoHandlers{service: mock}

	body, _ := json.Marshal(models.ToDoRequest{Title: "", Body: "Desc"})
	req := httptest.NewRequest("PUT", "/tasks/1", bytes.NewReader(body))
	req.SetPathValue("id", "1")
	rr := httptest.NewRecorder()

	handler.Update(rr, req)

	if rr.Code != 400 {
		t.Errorf("status = %d, want 400 for empty title", rr.Code)
	}
}
