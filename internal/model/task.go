package model

import "time"

// Task represents a to-do task.
// swagger:model
type Task struct {
    // Unique ID of the task
    ID string `json:"id"`
    // Title of the task
    Title string `json:"title"`
    // Status of the task: Pending, InProgress, Completed
    Status string `json:"status"`
    // Time when created
    CreatedAt time.Time `json:"created_at"`
    // Time when last updated
    UpdatedAt time.Time `json:"updated_at"`
}
