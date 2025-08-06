package repository

import (
    "errors"
    "sync"
    "time"

    "github.com/google/uuid"
    "github.com/vigneshv1095/task-service/internal/model"
)

// InMemoryRepo is an in-memory TaskRepository.
type InMemoryRepo struct {
    mu    sync.Mutex
    tasks map[string]*model.Task
}

// NewInMemoryRepo initializes an in-memory repo.
func NewInMemoryRepo() *InMemoryRepo {
    return &InMemoryRepo{tasks: make(map[string]*model.Task)}
}

func (r *InMemoryRepo) Create(task *model.Task) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    task.ID = uuid.New().String()
    task.CreatedAt = time.Now()
    task.UpdatedAt = time.Now()
    r.tasks[task.ID] = task
    return nil
}

func (r *InMemoryRepo) GetByID(id string) (*model.Task, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    if t, ok := r.tasks[id]; ok {
        return t, nil
    }
    return nil, errors.New("task not found")
}

func (r *InMemoryRepo) List(limit, offset int, status string) ([]*model.Task, int, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    var all []*model.Task
    for _, t := range r.tasks {
        if status != "" && t.Status != status {
            continue
        }
        all = append(all, t)
    }
    total := len(all)
    end := offset + limit
    if end > total {
        end = total
    }
    if offset > total {
        return []*model.Task{}, total, nil
    }
    return all[offset:end], total, nil
}

func (r *InMemoryRepo) Update(task *model.Task) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    if existing, ok := r.tasks[task.ID]; ok {
        existing.Title = task.Title
        existing.Status = task.Status
        existing.UpdatedAt = time.Now()
        return nil
    }
    return errors.New("task not found")
}

func (r *InMemoryRepo) Delete(id string) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, ok := r.tasks[id]; ok {
        delete(r.tasks, id)
        return nil
    }
    return errors.New("task not found")
}
