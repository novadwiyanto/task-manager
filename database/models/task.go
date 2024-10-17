package models

import (
	"task-manager/pkg/enum"
	"time"
)

type Status string

type Task struct {
	ID          uint          `json:"id" gorm:"primaryKey"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      enum.Status   `gorm:"type:ENUM('not started', 'in progress', 'completed');default:'not started'"`
	Priority    enum.Priority `gorm:"type:ENUM('low', 'medium', 'high');default:'low'"`
	DueDate     time.Time     `json:"due_date"`
	UserID      uint          `json:"user_id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
