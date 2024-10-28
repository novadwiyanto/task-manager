package models

import (
	"task-manager/pkg/enum"
	"time"
)

type TaskHistory struct {
	ID           int           `json:"id" gorm:"primaryKey"`
	TaskID       int           `json:"task_id"`
	Task         Task          `json:"task" gorm:"foreignKey:TaskID"`
	ChangedByID  int           `json:"changed_by_id"`
	ChangedBy    User          `json:"changed_by" gorm:"foreignKey:ChangedByID"`
	Description  string        `json:"description,omitempty" gorm:"type:text;null"`
	Status       enum.Status   `json:"status" gorm:"type:ENUM('not_started', 'in_progress', 'completed');default:'not_started'"`
	Priority     enum.Priority `json:"priority" gorm:"type:ENUM('low', 'medium', 'high');default:'low'"`
	DueDate      time.Time     `json:"due_date,omitempty" gorm:"null"`
	AssignedToID *uint         `json:"assigned_to_id,omitempty"`
	AssignedTo   *User         `json:"assigned_to,omitempty" gorm:"foreignKey:AssignedToID"`
	ChangedAt    time.Time     `json:"changed_at"`
}
