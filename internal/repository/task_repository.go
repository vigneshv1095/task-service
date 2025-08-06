package repository

import "github.com/vigneshv1095/task-service/internal/model"

// TaskRepository defines CRUD ops for tasks.
type TaskRepository interface {
    Create(task *model.Task) error
    GetByID(id string) (*model.Task, error)
    List(limit, offset int, status string) ([]*model.Task, int, error)
    Update(task *model.Task) error
    Delete(id string) error
}
