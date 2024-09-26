package request

import (
	"task-manager/pkg/enum"
	"time"
)

type CreateTaskRequest struct {
	Title string `validate:"required"`
	Description string
	Status enum.Status
	Priority enum.Priority
	DueDate time.Time
}

type UpdateTaskRequest struct {
	ID uint `validate:"required"`
	Title string `validate:"required"`
	Description string
	Status enum.Status
	Priority enum.Priority
	DueDate time.Time
}