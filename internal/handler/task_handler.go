package handler

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/vigneshv1095/task-service/internal/model"
    "github.com/vigneshv1095/task-service/internal/service"
    "github.com/vigneshv1095/task-service/pkg/pagination"
)

// TaskHandler handles HTTP requests for tasks.
type TaskHandler struct {
    svc *service.TaskService
}

// NewTaskHandler creates a new TaskHandler.
func NewTaskHandler(svc *service.TaskService) *TaskHandler {
    return &TaskHandler{svc: svc}
}

// Register routes.
func (h *TaskHandler) Register(r *mux.Router) {
    r.HandleFunc("/tasks", h.Create).Methods("POST")
    r.HandleFunc("/tasks", h.List).Methods("GET")
    r.HandleFunc("/tasks/{id}", h.Get).Methods("GET")
    r.HandleFunc("/tasks/{id}", h.Update).Methods("PUT")
    r.HandleFunc("/tasks/{id}", h.Delete).Methods("DELETE")
}

// Create a new task.
// @Summary Create task
// @Description Create a new task with title and status
// @Accept json
// @Produce json
// @Param task body model.Task true "Task payload"
// @Success 201 {object} model.Task
// @Router /tasks [post]
func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
    var t model.Task
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.svc.Create(&t); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(t)
}

// List tasks with pagination and optional status filter.
// @Summary List tasks
// @Description Get tasks with pagination and optional status
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param status query string false "Status filter"
// @Success 200 {object} map[string]interface{}
// @Router /tasks [get]
func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
    limit, offset := pagination.Parse(r)
    status := r.URL.Query().Get("status")

    tasks, total, err := h.svc.List(limit, offset, status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]interface{}{
        "items": tasks,
        "total": total,
    })
}

// Get a task by ID.
// @Summary Get task
// @Description Get task by ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} model.Task
// @Router /tasks/{id} [get]
func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    t, err := h.svc.Get(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(t)
}

// Update a task.
// @Summary Update task
// @Description Update task title and status
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body model.Task true "Updated Task"
// @Success 200 {object} model.Task
// @Router /tasks/{id} [put]
func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var t model.Task
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    t.ID = id
    if err := h.svc.Update(&t); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(t)
}

// Delete a task.
// @Summary Delete task
// @Description Delete task by ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 204 "No Content"
// @Router /tasks/{id} [delete]
func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    if err := h.svc.Delete(id); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
