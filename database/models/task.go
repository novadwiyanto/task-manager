package models

import (
	"task-manager/pkg/enum"
	"time"
)

type Task struct {
	ID           uint          `json:"id" gorm:"primaryKey"`
	Title        string        `json:"title" gorm:"type:varchar(255);not null"`
	Description  string        `json:"description,omitempty" gorm:"type:text;null"`
	Status       enum.Status   `json:"status" gorm:"type:ENUM('not_started', 'in_progress', 'completed');default:'not_started'"`
	Priority     enum.Priority `json:"priority" gorm:"type:ENUM('low', 'medium', 'high');default:'low'"`
	DueDate      time.Time     `json:"due_date,omitempty" gorm:"null"`
	CreatedByID  uint          `json:"created_by_id"`
	CreatedBy    User          `json:"created_by" gorm:"foreignKey:CreatedByID"`
	AssignedToID *uint         `json:"assigned_to_id,omitempty"`
	AssignedTo   *User         `json:"assigned_to,omitempty" gorm:"foreignKey:AssignedToID"`
	TeamID       *uint         `json:"team_id,omitempty"`
	Team         *Team         `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	CategoryID   *uint         `json:"category_id,omitempty"`
	Category     *Category     `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	CreatedAt    time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}
