package domain

type TaskRepository interface {
	GetByUserID(userID string) ([]*Task, error)
	GetByID(id string) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id string) error
}
