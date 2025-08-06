package service

import "github.com/vigneshv1095/task-service/internal/model"
import "github.com/vigneshv1095/task-service/internal/repository"

// TaskService handles business logic for tasks.
type TaskService struct {
    repo repository.TaskRepository
}

// NewTaskService creates a new TaskService.
func NewTaskService(r repository.TaskRepository) *TaskService {
    return &TaskService{repo: r}
}

func (s *TaskService) Create(t *model.Task) error {
    return s.repo.Create(t)
}

func (s *TaskService) Get(id string) (*model.Task, error) {
    return s.repo.GetByID(id)
}

func (s *TaskService) List(limit, offset int, status string) ([]*model.Task, int, error) {
    return s.repo.List(limit, offset, status)
}

func (s *TaskService) Update(t *model.Task) error {
    return s.repo.Update(t)
}

func (s *TaskService) Delete(id string) error {
    return s.repo.Delete(id)
}
