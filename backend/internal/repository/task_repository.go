package postgres

import (
	"database/sql"
	"time"

	"github.com/tsunakit99/sp-web2/backend/internal/domain"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) domain.TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetByUserID(userID string) ([]*domain.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, title, description, is_completed, due_date, created_at
		FROM tasks WHERE user_id = $1 ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		var t domain.Task
		err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.IsCompleted, &t.DueDate, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

func (r *taskRepository) GetByID(id string) (*domain.Task, error) {
	row := r.db.QueryRow(`
		SELECT id, user_id, title, description, is_completed, due_date, created_at
		FROM tasks WHERE id = $1
	`, id)

	var t domain.Task
	err := row.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.IsCompleted, &t.DueDate, &t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *taskRepository) Create(task *domain.Task) error {
	query := `
		INSERT INTO tasks (user_id, title, description, is_completed, due_date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	// `created_at` は自前で指定しない場合、SQL側で `default now()` が動く

	err := r.db.QueryRow(
		query,
		task.UserID,
		task.Title,
		task.Description,
		task.IsCompleted,
		task.DueDate,
		time.Now(), // 任意。DBに任せたいならここも省略できる
	).Scan(&task.ID, &task.CreatedAt)

	return err
}


func (r *taskRepository) Update(task *domain.Task) error {
	query := `
		UPDATE tasks SET title = $1, description = $2, is_completed = $3, due_date = $4
		WHERE id = $5
	`
	_, err := r.db.Exec(query, task.Title, task.Description, task.IsCompleted, task.DueDate, task.ID)
	return err
}

func (r *taskRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	return err
}
