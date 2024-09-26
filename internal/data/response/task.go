package response

import (
	"task-manager/pkg/enum"
	"time"
)

type Task struct {
	ID          uint
	Title       string
	Description string
	Status      enum.Status
	Priority    enum.Priority
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
