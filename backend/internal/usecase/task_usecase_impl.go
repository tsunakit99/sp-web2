package usecase

import "github.com/tsunakit99/sp-web2/backend/internal/domain"

type taskUsecase struct {
	repo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase) GetTasks(userID string) ([]*domain.Task, error) {
	return u.repo.GetByUserID(userID)
}

func (u *taskUsecase) GetTask(id string) (*domain.Task, error) {
	return u.repo.GetByID(id)
}

func (u *taskUsecase) CreateTask(task *domain.Task) error {
	return u.repo.Create(task)
}

func (u *taskUsecase) UpdateTask(task *domain.Task) error {
	return u.repo.Update(task)
}

func (u *taskUsecase) DeleteTask(id string) error {
	return u.repo.Delete(id)
}
