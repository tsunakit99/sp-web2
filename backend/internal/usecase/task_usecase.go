package usecase

import "github.com/tsunakit99/sp-web2/backend/internal/domain"

type TaskUsecase interface {
	GetTasks(userID string) ([]*domain.Task, error)
	GetTask(id string) (*domain.Task, error)
	CreateTask(task *domain.Task) error
	UpdateTask(task *domain.Task) error
	DeleteTask(id string) error
}
